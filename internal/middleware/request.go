package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成新的请求 ID
		requestID := uuid.New().String()

		// 将请求 ID 设置到上下文
		c.Set("X-Request-ID", requestID)

		// 将请求 ID 添加到响应头
		c.Writer.Header().Set("X-Request-ID", requestID)

		// 继续处理请求
		c.Next()
	}
}
