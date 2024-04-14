/**
 * @Author: chentong
 * @Date: 2023/11/23 10:34
 */

package xdatatype

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// mysql json 类型存储 ["a", "b"]
type StringArray []string

// mysql json 类型存储 [1, 2]
type Int64Array []int64

// Value return json value, implement driver.Valuer interface
func (j *StringArray) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	result, err := json.Marshal(j)
	if err != nil {
		return nil, nil
	}
	return string(result), nil
}

// Scan value into Jsonb, implements sql.Scanner interface
func (j *StringArray) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		if len(v) > 0 {
			bytes = make([]byte, len(v))
			copy(bytes, v)
		}
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var s []string
	_ = json.Unmarshal(bytes, &s)

	*j = s
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Int64Array) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	result, err := json.Marshal(j)
	if err != nil {
		return nil, nil
	}
	return string(result), nil
}

// Scan value into Jsonb, implements sql.Scanner interface
func (j *Int64Array) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		if len(v) > 0 {
			bytes = make([]byte, len(v))
			copy(bytes, v)
		}
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var s []int64
	_ = json.Unmarshal(bytes, &s)

	*j = s
	return nil
}
