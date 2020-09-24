package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mogfee/blog-server/global"
	"github.com/mogfee/blog-server/pkg/app"
	"github.com/mogfee/blog-server/pkg/errcode"
	"github.com/mogfee/blog-server/pkg/xmail"
	"github.com/xinliangnote/go-util/time"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utilGin := app.NewGinWap(c)
				span, _ := utilGin.StartSpanFromContext("异常处理")
				defer span.Finish()
				span.SetTag("error", fmt.Sprintf("%s", err))

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				traceId, _ := c.Get("X-Trace-Id")
				subject := fmt.Sprintf("【重要错误】%v %s 项目出错了！", traceId, global.AppSetting.ServerName)
				body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				body = strings.ReplaceAll(body, "{RequestTime}", time.GetCurrentDate())
				body = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI)
				body = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
				body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)
				options := &xmail.Options{
					MailHost: global.MonitorSetting.SystemEmailHost,
					MailPort: global.MonitorSetting.SystemEmailPort,
					MailUser: global.MonitorSetting.SystemEmailUser,
					MailPass: global.MonitorSetting.SystemEmailPass,
					MailTo:   global.MonitorSetting.ErrorNotifyUser,
					Subject:  subject,
					Body:     body,
				}
				if global.ServerSetting.RunModel != "debug" {
					_ = xmail.Send(options)
				}
				utilGin.ToErrorResponse(errcode.ServerError)
			}
		}()
		c.Next()
	}
}
