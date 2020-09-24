package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/mogfee/blog-server/global"
	"github.com/mogfee/blog-server/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(body []byte) (int, error) {
	if n, err := w.body.Write(body); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(body)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter
		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		traceId, _ := c.Get("X-Trace-Id")
		spanId, _ := c.Get("X-Span-Id")
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
			"trace_id": traceId,
			"span_id":  spanId,
		}
		global.Logger.WithFields(fields).
			Infof("access_log: method: %s,status_code:%d, begin_time: %d, end_time:%d",
				c.Request.Method,
				bodyWriter.Status(),
				beginTime,
				endTime,
			)
	}
}
