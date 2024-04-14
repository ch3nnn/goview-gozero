package config

import (
	"context"
	"errors"
	"testing"
	"time"

	"gorm.io/gorm/logger"
)

func TestNewDatabaseLogger(t *testing.T) {
	l := NewDatabaseLogger(logger.Info, 200*time.Millisecond)
	l.Info(context.Background(), "info test, p1: %v, p2: %v, p3: %v", 1, 2, 3)
	l.Warn(context.Background(), "warn test, p1: %v, p2: %v, p3: %v", 1, 2, 3)
	l.Error(context.Background(), "error test, p1: %v, p2: %v, p3: %v", 1, 2, 3)
	l.Trace(context.Background(), time.Now(), func() (string, int64) { return "test sql", 1 }, nil)
	l.Trace(context.Background(), time.Now(), func() (string, int64) { return "test err sql", 0 }, errors.New("err sql"))

	l = NewDatabaseLogger(logger.Error, 200*time.Millisecond)
	l.Info(context.Background(), "info test, p1: %v, p2: %v, p3: %v", 1, 2, 3)
	l.Warn(context.Background(), "warn test, p1: %v, p2: %v, p3: %v", 1, 2, 3)
	l.Error(context.Background(), "error test, p1: %v, p2: %v, p3: %v", 1, 2, 3)
	l.Trace(context.Background(), time.Now(), func() (string, int64) { return "test sql", 1 }, nil)
	l.Trace(context.Background(), time.Now(), func() (string, int64) { return "test err sql", 0 }, errors.New("err sql"))
}
