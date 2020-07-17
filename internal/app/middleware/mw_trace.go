package middleware

import (
	"github.com/tanjiancheng/gin-amis-admin/internal/app/icontext"
	"github.com/tanjiancheng/gin-amis-admin/pkg/logger"
	"github.com/tanjiancheng/gin-amis-admin/pkg/trace"
	"github.com/gin-gonic/gin"
)

// TraceMiddleware 跟踪ID中间件
func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		// 优先从请求头中获取请求ID
		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = trace.NewID()
		}

		ctx := icontext.NewTraceID(c.Request.Context(), traceID)
		ctx = logger.NewTraceIDContext(ctx, traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", traceID)

		c.Next()
	}
}
