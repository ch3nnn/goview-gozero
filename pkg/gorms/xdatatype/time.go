/**
 * @Author: chentong
 * @Date: 2023/11/23 10:31
 */

package xdatatype

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// time 转 int64
// 毫秒时间戳
type TimeToInt64 int64

// Value implements driver.Valuer interface and returns string format of Time.
func (t *TimeToInt64) Value() (driver.Value, error) {
	return time.UnixMilli(int64(*t)), nil
}

// Scan implements sql.Scanner interface and scans value into Time,
func (t *TimeToInt64) Scan(src interface{}) error {
	switch v := src.(type) {
	case time.Time:
		*t = TimeToInt64(v.UnixMilli())
	case string: // 时间字符串类型
		parseTime, _ := time.Parse(time.DateTime, v)
		*t = TimeToInt64(parseTime.UnixMilli())
	case []uint8:
		loc, _ := time.LoadLocation("Asia/Shanghai")
		parseTime, _ := time.ParseInLocation(time.DateTime, string(v), loc)
		*t = TimeToInt64(parseTime.UnixMilli())
	default:
		return errors.New(fmt.Sprintf("failed to scan value: %v", v))
	}

	return nil
}
