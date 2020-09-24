package model

import (
	"fmt"
	"github.com/mogfee/blog-server/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username, databaseSetting.Password, databaseSetting.Host,
		databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: NewLogger(log.New(os.Stdout, "\r\n", log.LstdFlags), Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      true,
		}),
	})
	//
	//
	//db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//if global.ServerSetting.RunModel == "debug" {
	//	db.LogMode(true)
	//}
	////
	////db.SingularTable(true)
	//db.DB().SetConnMaxLifetime(databaseSetting.ConnMaxLifetime)
	////db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	//db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}
