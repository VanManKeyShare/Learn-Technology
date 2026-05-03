package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RateLimitConfig struct {
	Limit     int
	Window    time.Duration
	KeyPrefix string
	SkipOnErr bool // NẾU REDIS LỖI THÌ CHO PASS QUA (FAIL-OPEN)
}

type RateLimiter struct {
	rdb    *redis.Client
	config RateLimitConfig
	script *redis.Script
}

// LUA SCRIPT: ATOMIC CHECK-AND-INCREMENT
// TRẢ VỀ [CURRENT_COUNT, TTL_MS]
var rateLimitLua = redis.NewScript(`
local key     = KEYS[1]
local limit   = tonumber(ARGV[1])
local window  = tonumber(ARGV[2])  -- milliseconds

local current = redis.call("INCR", key)
if current == 1 then
    redis.call("PEXPIRE", key, window)
end

local ttl = redis.call("PTTL", key)
return {current, ttl}
`)

func NewRateLimiter(rdb *redis.Client, cfg RateLimitConfig) *RateLimiter {
	if cfg.KeyPrefix == "" {
		cfg.KeyPrefix = "rlm-global"
	}
	return &RateLimiter{rdb: rdb, config: cfg, script: rateLimitLua}
}

type result struct {
	count     int
	remaining int
	resetMs   int64
	allowed   bool
}

func (rl *RateLimiter) check(ctx context.Context, ip string) (*result, error) {
	key := fmt.Sprintf("%s:%s", rl.config.KeyPrefix, ip)
	windowMs := rl.config.Window.Milliseconds()

	vals, err := rl.script.Run(ctx, rl.rdb,
		[]string{key},
		rl.config.Limit,
		windowMs,
	).Int64Slice()
	if err != nil {
		return nil, fmt.Errorf("REDIS EVAL: %w", err)
	}

	count := int(vals[0])
	ttlMs := vals[1]

	return &result{
		count:     count,
		remaining: max(0, rl.config.Limit-count),
		resetMs:   ttlMs,
		allowed:   count <= rl.config.Limit,
	}, nil
}

func (rl *RateLimiter) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		res, err := rl.check(c.Request.Context(), ip)

		if err != nil {
			// FAIL-OPEN HOẶC FAIL-CLOSED TUỲ CONFIG
			if rl.config.SkipOnErr {
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "RATE LIMITER UNAVAILABLE",
			})
			return
		}

		// GẮN HEADERS CHUẨN RFC 6585
		c.Header("X-RateLimit-Limit", strconv.Itoa(rl.config.Limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(res.remaining))
		c.Header("X-RateLimit-Reset", strconv.FormatInt(res.resetMs/1000, 10))

		if !res.allowed {
			retryAfter := res.resetMs / 1000
			if retryAfter < 1 {
				retryAfter = 1
			}

			// For Debug
			// c.Header("Rate-Limit-Key", rl.config.KeyPrefix)

			c.Header("Retry-After", strconv.FormatInt(retryAfter, 10))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":       "TOO MANY REQUESTS",
				"retry_after": retryAfter,
			})
			return
		}

		c.Next()
	}
}
