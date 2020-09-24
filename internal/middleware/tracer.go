package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mogfee/blog-server/global"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var ctx context.Context

		span := opentracing.SpanFromContext(c.Request.Context())
		if span != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), global.Tracer, c.Request.URL.Path)
		}
		defer span.Finish()

		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			c.Set("X-Trace-Id", sc.TraceID().String())
			c.Set("X-Span-Id", sc.SpanID().String())
			c.Header("X-Request-id", sc.TraceID().String())
		}
		c.Set("tracer_span", span)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
