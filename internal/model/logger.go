package model

import (
	"context"
	"github.com/mogfee/blog-server/global"
	logger2 "github.com/mogfee/blog-server/pkg/logger"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"log"
	"time"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// Writer log writer interface
type Writer interface {
	Printf(string, ...interface{})
}
type MyWrite struct {
}

func (m MyWrite) Printf(msg string, data ...interface{}) {
	log.Printf(msg, data...)
}

type Config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      logger.LogLevel
}

func NewLogger(writer Writer, config Config) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%d] %s"
		traceWarnStr = "%s\n[%.3fms] [rows:%d] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%d] %s"
	)

	if config.Colorful {
		infoStr = Green + "%s\n" + Reset + Green + "[info] " + Reset
		warnStr = Blue + "%s\n" + Reset + Magenta + "[warn] " + Reset
		errStr = Magenta + "%s\n" + Reset + Red + "[error] " + Reset
		traceStr = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + Blue + "[rows:%d]" + Reset + " %s"
		traceWarnStr = Green + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%d]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + Blue + "[rows:%d]" + Reset + " %s"
	}

	return &mylogger{
		Writer:       writer,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type mylogger struct {
	Writer
	Config
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func (l *mylogger) LogMode(level logger.LogLevel) logger.Interface {
	newmylogger := *l
	newmylogger.LogLevel = level
	return &newmylogger
}

func (l *mylogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l *mylogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l *mylogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l *mylogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		sql, rows := fc()
		if err != nil {
			global.Logger.WithContext(ctx).WithFields(logger2.Fields{
				"error": err.Error(),
			}).Error(sql)
		} else {
			global.Logger.WithContext(ctx).Debug(sql)
		}
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			l.Printf(l.traceWarnStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case l.LogLevel >= logger.Info:
			l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
