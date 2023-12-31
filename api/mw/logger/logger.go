package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func TraceLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 每个请求生成的请求traceId具有全局唯一性
		u1, _ := uuid.NewUUID()
		traceId := u1.String()
		NewContext(ctx, zap.String("traceId", traceId))

		// 为日志添加请求的地址以及请求参数等信息
		NewContext(ctx, zap.String("request.method", ctx.Request.Method))

		headers, _ := json.Marshal(ctx.Request.Header)
		NewContext(ctx, zap.String("request.headers", string(headers)))

		NewContext(ctx, zap.String("request.url", ctx.Request.URL.String()))
		// 将请求参数json序列化后添加进日志上下文
		if ctx.Request.Form == nil {
			err := ctx.Request.ParseMultipartForm(32 << 20)
			if err != nil {
				return
			}
		}
		form, _ := json.Marshal(ctx.Request.Form)
		NewContext(ctx, zap.String("request.params", string(form)))
		ctx.Next()
	}
}
