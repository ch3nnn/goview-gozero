package config

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-stack/stack"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

const (
	traceInfo  = "%s [%.3fms] [rows:%v] %s"
	traceWarn  = "%s %s [%.3fms] [rows:%v] %s"
	traceError = "%s %s [%.3fms] [rows:%v] %s"
)

// DatabaseLogger 日志记录器
type DatabaseLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

// NewDatabaseLogger 新建日志记录器
func NewDatabaseLogger(logLevel logger.LogLevel, slowThreshold time.Duration) *DatabaseLogger {
	return &DatabaseLogger{LogLevel: logLevel, SlowThreshold: slowThreshold}
}

// LogMode 设置日志记录模式
func (l *DatabaseLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info Info日志记录
func (l *DatabaseLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		logx.WithContext(ctx).Infof(msg, data...)
	}
}

// Warn Warn日志记录
func (l *DatabaseLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		logx.WithContext(ctx).Slowf(msg, data...)
	}
}

// Error Error日志记录
func (l *DatabaseLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		logx.WithContext(ctx).Errorf(msg, data...)
	}
}

// Trace Trace日志记录
func (l *DatabaseLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > logger.Silent {
		log := logx.WithContext(ctx)
		elapsed := time.Since(begin)

		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				log.Errorf(traceError, FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				log.Errorf(traceError, FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("Slow SQL Greater Than %v", l.SlowThreshold)
			if rows == -1 {
				log.Slowf(traceWarn, FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				log.Slowf(traceWarn, FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel == logger.Info:
			sql, rows := fc()
			if rows == -1 {
				log.Infof(traceInfo, FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				log.Infof(traceInfo, FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

// FileWithLineNum 获取调用堆栈信息
func FileWithLineNum() string {
	cs := stack.Trace().TrimBelow(stack.Caller(2)).TrimRuntime()

	for _, c := range cs {
		s := fmt.Sprintf("%+v", c)
		if !strings.Contains(s, "gorm.io/") && !strings.Contains(s, ".gen.go") {
			return s
		}
	}

	return ""
}
