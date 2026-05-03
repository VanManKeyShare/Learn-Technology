package main

import (
	"log"
	"os"
	"strings"
	"time"
	Handler "vmk-gin-app-docker/module/handler"
	"vmk-gin-app-docker/module/middleware"
	"vmk-gin-app-docker/module/response"
	MySQL "vmk-gin-app-docker/module/service/db"
	Redis "vmk-gin-app-docker/module/service/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	/*
		SET GIN TO RELEASE MODE TO DISABLE DEBUG OUTPUT
		THIS IS IMPORTANT WHEN RUNNING IN A PRODUCTION ENVIRONMENT

		gin.SetMode(gin.ReleaseMode)
	*/

	// INIT MYSQL
	if err := MySQL.Init_MySQL_DB(); err != nil {
		log.Fatalf("DATABASE INIT FAILED: %v", err)
	}
	defer MySQL.DB.Close()

	// INIT REDIS
	if err := Redis.Init_Redis(); err != nil {
		log.Fatalf("REDIS INIT FAILED: %v", err)
	}
	defer Redis.RedisClient.Close()

	// INIT GIN
	r := gin.Default()

	// CẤU HÌNH TIN TƯỞNG CÁC IP REVERSE PROXY (NHƯ NGINX, CLOUDFLARE, HOẶC AWS ALB) ĐỂ PHÂN GIẢI CLIENT IP
	proxies := os.Getenv("TRUSTED_PROXIES")
	if proxies != "" {
		r.SetTrustedProxies(strings.Split(proxies, ","))
	}

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://example.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "X-RateLimit-Limit", "X-RateLimit-Remaining", "X-RateLimit-Reset"},
		AllowCredentials: true, // CHO PHÉP GỬI COOKIE, KHÔNG SỬ DỤNG KHI ORIGIN CÓ WILDCARD (*)
		MaxAge:           12 * time.Hour,
	}))

	// RATE LIMIT MIDDLEWARE (60 REQUEST MỖI PHÚT)
	// ── GLOBAL LIMIT ──
	globalRL := middleware.NewRateLimiter(Redis.RedisClient, middleware.RateLimitConfig{
		Limit:     60,
		Window:    time.Minute,
		KeyPrefix: "rlm-global",
		SkipOnErr: true, // PRODUCTION: FAIL-OPEN ĐỂ KHÔNG BLOCK USER KHI REDIS RESTART
	})
	r.Use(globalRL.Handler())

	// ── STRICT LIMIT CHO AUTH ROUTES ──
	authRL := middleware.NewRateLimiter(Redis.RedisClient, middleware.RateLimitConfig{
		Limit:     10,
		Window:    time.Minute,
		KeyPrefix: "rlm-auth",
		SkipOnErr: false, // AUTH PHẢI CHẶT: FAIL-CLOSED
	})
	auth := r.Group("/auth")
	auth.Use(authRL.Handler())
	{
		auth.GET("/login", func(c *gin.Context) {
			response.Success(c, "OK", "Login Success")
		})
		auth.POST("/register", func(c *gin.Context) {
			response.Success(c, "OK", "Register Success")
		})
	}

	r.GET("/", func(c *gin.Context) {
		response.Success(c, "Hello Gin Docker", "OK")
	})

	r.GET("/db", Handler.Check_Health_Database)

	r.GET("/redis", Handler.Test_Redis_Counter)

	r.Run(":8080")
}
