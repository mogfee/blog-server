package global

import (
	"github.com/mogfee/blog-server/pkg/logger"
	"github.com/mogfee/blog-server/pkg/setting"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServiceSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JaegerSetting *setting.JaegerSettingS
	MonitorSetting *setting.MonitorSettingS

	DBEngine *gorm.DB

	Logger *logger.Logger
	Tracer opentracing.Tracer
)
