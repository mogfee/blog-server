package model

import (
	"fmt"
	"github.com/mogfee/blog-server/global"
	logger2 "github.com/mogfee/blog-server/pkg/logger"
	"github.com/mogfee/blog-server/pkg/setting"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

func TestNewDBEngine(t *testing.T) {
	global.Logger = logger2.NewLogger(os.Stdout, "", 1)

	db, err := NewDBEngine(&setting.DatabaseSettingS{
		DBType:    "mysql",
		Username:  "root",
		Password:  "123456",
		Host:      "localhost:3306",
		DBName:    "article",
		Charset:   "utf8mb4",
		ParseTime: true,
	})
	if err != nil {
		panic(err)
	}
	db.Logger = NewLogger(log.New(os.Stdout, "\r\n", log.LstdFlags), Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      false,
	})
	fmt.Println(db.Table("123123").Where("id=1").Updates(map[string]interface{}{
		"123": 123,
	}).Error)
}
