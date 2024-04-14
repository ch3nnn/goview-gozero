package config

import "gorm.io/gorm/logger"

type DBType string

const (
	Mysql  DBType = "mysql"
	Pgsql  DBType = "postgres"
	Sqlite DBType = "sqlite3"
	NoType DBType = "none"
)

type DBLogLevel string

const (
	Info   DBLogLevel = "info"
	Warn   DBLogLevel = "warn"
	Error  DBLogLevel = "error"
	Silent DBLogLevel = "silent"
)

func (l DBLogLevel) ToGormLogLevel() (logLevel logger.LogLevel) {
	switch l {
	case Info:
		logLevel = logger.Info
	case Warn:
		logLevel = logger.Warn
	case Error:
		logLevel = logger.Error
	case Silent:
		logLevel = logger.Silent
	}

	return
}
