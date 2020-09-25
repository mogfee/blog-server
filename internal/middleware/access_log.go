package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mogfee/blog-server/global"
	"github.com/mogfee/blog-server/pkg/logger"
	"github.com/uber/jaeger-client-go"
	"io/ioutil"
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

		body, err := c.GetRawData()
		var postBody string
		if err == nil {
			postBody = fmt.Sprintf("%s", body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}

		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter
		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		spanContext, ok := c.Get("tracer_span")
		if ok {
			span := spanContext.(*jaeger.Span)
			if span != nil {
				if bodyWriter.Status() != 200 {
					span.SetTag("error", true)
				}
				span.SetBaggageItem("post", postBody)
				span.SetBaggageItem("ip", c.ClientIP())
				span.SetTag("http.status_code", bodyWriter.Status())
			}
		}
		traceId, _ := c.Get("X-Trace-Id")
		spanId, _ := c.Get("X-Span-Id")
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
			"trace_id": traceId,
			"span_id":  spanId,
		}
		global.Logger.WithCaller(1).WithFields(fields).
			Infof("access_log: url: %s, method: %s, body:%s, status_code:%d, begin_time: %d, end_time:%d",
				c.Request.URL,
				c.Request.Method,
				postBody,
				bodyWriter.Status(),
				beginTime,
				endTime,
			)
	}
}
