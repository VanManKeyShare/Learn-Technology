package handler

import (
	"context"
	"time"
	"vmk-gin-app-docker/module/response"
	Redis "vmk-gin-app-docker/module/service/db"

	"github.com/gin-gonic/gin"
)

func Test_Redis_Counter(c *gin.Context) {

	// KIỂM TRA KẾT NỐI
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Redis.RedisClient.Ping(ctx).Err(); err != nil {
		response.InternalError(c, "FAILED TO CONNECT TO REDIS", err.Error())
		return
	}

	// DÙNG INCR TRỰC TIẾP — ATOMIC, KHÔNG RACE CONDITION
	// NẾU KEY CHƯA TỒN TẠI, REDIS TỰ KHỞI TẠO = 0 RỒI TĂNG LÊN 1
	NewCounter, err := Redis.RedisClient.Incr(ctx, "counter").Result()
	if err != nil {
		response.InternalError(c, "FAILED TO INCREMENT COUNTER", err.Error())
		return
	}

	response.Success(c, "OK", gin.H{
		"counter": NewCounter,
	})
}
