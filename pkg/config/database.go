package config

import (
	"fmt"
	"os"
	"time"

	"github.com/Pacific73/gorm-cache/cache"
	"github.com/Pacific73/gorm-cache/config"
	redis "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	redis2 "github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConf stores database configurations.
type DatabaseConf struct {
	Host          string     `json:",env=DATABASE_HOST"`
	Port          int        `json:",env=DATABASE_PORT"`
	Username      string     `json:",default=root,env=DATABASE_USERNAME"`
	Password      string     `json:",optional,env=DATABASE_PASSWORD"`
	DBName        string     `json:",default=sa,env=DATABASE_DBNAME"`
	SSLMode       string     `json:",optional,env=DATABASE_SSL_MODE"`
	Type          DBType     `json:",default=mysql,options=[mysql,postgres,sqlite3],env=DATABASE_TYPE"`
	MaxOpenConn   int        `json:",optional,default=100,env=DATABASE_MAX_OPEN_CONN"`
	CacheTime     int        `json:",optional,default=10,env=DATABASE_CACHE_TIME"`
	DBPath        string     `json:",optional,env=DATABASE_DBPATH"`
	MysqlConfig   string     `json:",optional,env=DATABASE_MYSQL_CONFIG"`
	PGConfig      string     `json:",optional,env=DATABASE_PG_CONFIG"`
	SqliteConfig  string     `json:",optional,env=DATABASE_SQLITE_CONFIG"`
	LogLevel      DBLogLevel `json:",default=info,options=[info,warn,error,silent],env=DATABASE_LOG_LEVEL"` // 日志级别，枚举（info、warn、error和silent）
	SlowThreshold int        `json:",optional,default=1000"`                                                // 记录慢查询阈值，单位毫秒
}

// Deprecated: NewCacheDriver returns a Gorm driver with cache.
// 不太稳定，不建议使用
func (c DatabaseConf) NewCacheDriver(redisConf redis2.RedisConf) *gorm.DB {
	db := c.NewDriver()

	rdb := redis.NewClient(&redis.Options{Addr: redisConf.Host, Password: redisConf.Pass})

	gorm2Cache, _ := cache.NewGorm2Cache(&config.CacheConfig{
		CacheLevel:           config.CacheLevelAll,
		CacheStorage:         config.CacheStorageRedis,
		RedisConfig:          cache.NewRedisConfigWithClient(rdb),
		InvalidateWhenUpdate: true, // when you create/update/delete objects, invalidate cache
		CacheTTL:             5000, // 5000 ms
		CacheMaxItemCnt:      5,    // if length of objects retrieved one single time
		// exceeds this number, then don't cache
	})

	err := db.Use(gorm2Cache)
	logx.Must(err)

	return db
}

// NewDriver  returns a Gorm driver without cache.
func (c DatabaseConf) NewDriver() (db *gorm.DB) {
	db = c.DBOpen(&gorm.Config{
		Logger:      c.GetLogger(),
		QueryFields: true,
	})

	// Connection Pool
	var err error
	if sqlDB, err := db.DB(); err != nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(c.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	logx.Must(err)

	return
}

// MysqlDSN returns mysql DSN.
func (c DatabaseConf) MysqlDSN() string {
	if c.MysqlConfig == "" {
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&timeout=5s", c.Username, c.Password, c.Host, c.Port, c.DBName)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.DBName, c.MysqlConfig)
}

// PostgresDSN returns Postgres DSN.
func (c DatabaseConf) PostgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", c.Host, c.Username, c.Password, c.DBName, c.Port)
}

// SqliteDSN returns Sqlite DSN.
func (c DatabaseConf) SqliteDSN() string {
	if c.DBPath == "" {
		logx.Must(errors.New("the database file path cannot be empty"))
	}

	if _, err := os.Stat(c.DBPath); os.IsNotExist(err) {
		f, err := os.OpenFile(c.DBPath, os.O_CREATE|os.O_RDWR, 0o600)
		if err != nil {
			logx.Must(fmt.Errorf("failed to create SQLite database file %q", c.DBPath))
		}
		if err := f.Close(); err != nil {
			logx.Must(fmt.Errorf("failed to create SQLite database file %q", c.DBPath))
		}
	} else {
		if err := os.Chmod(c.DBPath, 0o660); err != nil {
			logx.Must(fmt.Errorf("unable to set permission code on %s: %v", c.DBPath, err))
		}
	}

	return fmt.Sprintf("file:%s?_busy_timeout=100000&_fk=1%s", c.DBPath, c.SqliteConfig)
}

// GetDSN returns DSN according to the database type.
func (c DatabaseConf) GetDSN() string {
	switch c.Type {
	case Mysql:
		return c.MysqlDSN()
	case Pgsql:
		return c.PostgresDSN()
	case Sqlite:
		return c.SqliteDSN()
	}
	return string(NoType)
}

// Open initialize db session based on dialer
func (c DatabaseConf) DBOpen(opts ...gorm.Option) (db *gorm.DB) {
	var err error
	var dialer gorm.Dialector

	switch c.Type {
	case Mysql:
		dialer = mysql.Open(c.GetDSN())
	case Pgsql:
		dialer = postgres.Open(c.GetDSN())
	case Sqlite:
		dialer = sqlite.Open(c.GetDSN())
	default:
		dialer = mysql.Open(c.GetDSN())
	}

	db, err = gorm.Open(dialer, opts...)
	logx.Must(err)

	return
}

func (c DatabaseConf) GetLogger() logger.Interface {
	return NewDatabaseLogger(c.LogLevel.ToGormLogLevel(), time.Millisecond*time.Duration(c.SlowThreshold))
}
