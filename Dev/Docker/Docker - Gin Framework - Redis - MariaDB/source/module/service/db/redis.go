package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init_Redis() error {

	// LẤY THÔNG SỐ TỪ ENVIRONMENT
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	// VALIDATE ENV VARS
	if host == "" || port == "" {
		return fmt.Errorf("MISSING REQUIRED REDIS ENVIRONMENT VARIABLES")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,

		// CONNECTION POOL
		PoolSize:        10,
		MinIdleConns:    3,
		ConnMaxLifetime: 5 * time.Minute,
		ConnMaxIdleTime: 2 * time.Minute,

		// TIMEOUT
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// KIỂM TRA KẾT NỐI
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("FAILED TO CONNECT TO REDIS: %w", err)
	}

	fmt.Println("REDIS CONNECTED:", host, port)
	return nil
}
