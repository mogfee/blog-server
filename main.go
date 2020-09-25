package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mogfee/blog-server/global"
	"github.com/mogfee/blog-server/internal/model"
	"github.com/mogfee/blog-server/internal/routers"
	"github.com/mogfee/blog-server/pkg/logger"
	"github.com/mogfee/blog-server/pkg/setting"
	"github.com/mogfee/blog-server/pkg/tracer"
	"github.com/opentracing/opentracing-go"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}

	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}

	if err := setupTracer(); err != nil {
		log.Fatalf("init.setupTracer err:%v", err)
	}
}
func main() {
	gin.SetMode(global.ServerSetting.RunModel)
	router := routers.NewRouter()
	s := http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
func setupSetting() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err := settings.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err := settings.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err := settings.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	if err := settings.ReadSection("Jaeger", &global.JaegerSetting); err != nil {
		return err
	}
	if err := settings.ReadSection("Monitor", &global.MonitorSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.DatabaseSetting.ConnMaxLifetime *= time.Second

	return nil
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
func setupLogger() error {
	var out io.Writer
	if global.ServerSetting.RunModel != "debug" {
		out = &lumberjack.Logger{
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		}
	} else {
		out = os.Stdout
	}
	global.Logger = logger.NewLogger(out, "", log.LstdFlags)
	return nil
}
func setupTracer() error {
	tracer, _, err := tracer.NewJaegerTracer(global.AppSetting.ServerName, fmt.Sprintf("%s:%d", global.JaegerSetting.Host, global.JaegerSetting.Port))
	if err != nil {
		return err
	}
	opentracing.SetGlobalTracer(tracer)
	global.Tracer = tracer
	return nil
}
