package plover

import (
	"net/url"
	`strconv`
	"strings"
)

type Query struct {
	url.Values
}

func (q Query) HasQuery() bool {
	return q.Values != nil
}
func (q Query) Has(key string) bool {
	return !strings.EqualFold(q.Values.Get(key), "")
}
func (q Query) Get(key string, defValue ...interface{}) interface{} {
	var defaultValue interface{}
	
	if len(defValue) > 0 {
		defaultValue = defValue[0]
	} else {
		defaultValue = nil
	}
	
	if q.Values == nil {
		return defaultValue
	}
	vs := q.Values[key]
	if len(vs) == 0 {
		return defaultValue
	}
	return vs[0]
}
func (q Query) GetFloat32(key string, defValue ...float32) float32 {
	var defaultValue float32
	if len(defValue) > 0 {
		defaultValue = defValue[0]
	} else {
		defaultValue = 0
	}
	if q.Values == nil {
		return defaultValue
	}
	vs := q.Values[key]
	if len(vs) == 0 {
		return defaultValue
	}
	value, _ := strconv.ParseFloat(vs[0], 32)
	return float32(value)
}
func (q Query) GetFloat64(key string, defValue ...float64) float64 {
	var defaultValue float64
	if len(defValue) > 0 {
		defaultValue = defValue[0]
	} else {
		defaultValue = 0
	}
	if q.Values == nil {
		return defaultValue
	}
	vs := q.Values[key]
	if len(vs) == 0 {
		return defaultValue
	}
	value, _ := strconv.ParseFloat(vs[0], 32)
	return value
}
// 获取int数据
func (q Query) GetInt(key string, defValue ...int) int {
	var defaultValue int
	if len(defValue) > 0 {
		defaultValue = defValue[0]
	} else {
		defaultValue = 0
	}
	if q.Values == nil {
		return defaultValue
	}
	vs := q.Values[key]
	if len(vs) == 0 {
		return defaultValue
	}
	value, _ := strconv.ParseInt(vs[0], 10, 64)
	return int(value)
}
// 获取int64数据
func (q Query) GetInt64(key string, defValue ...int64) int64 {
	var defaultValue int64
	if len(defValue) > 0 {
		defaultValue = defValue[0]
	} else {
		defaultValue = 0
	}
	if q.Values == nil {
		return defaultValue
	}
	vs := q.Values[key]
	if len(vs) == 0 {
		return defaultValue
	}
	value, _ := strconv.ParseInt(vs[0], 10, 64)
	return value
}
// 获取int64数据
func (q Query) GetString(key string, defValue ...string) string {
	var defaultValue string
	if len(defValue) > 0 {
		defaultValue = defValue[0]
	} else {
		defaultValue = ""
	}
	if q.Values == nil {
		return defaultValue
	}
	vs := q.Values[key]
	if len(vs) == 0 {
		return defaultValue
	}
	return vs[0]
}
func (q Query) Set(key, value string) {
	q.Values.Set(key, value)
}

// Add adds the value to key. It appends to any existing
// values associated with key.
func (q Query) Add(key, value string) {
	q.Values.Add(key, value)
}

// Del deletes the values associated with key.
func (q Query) Del(key string) {
	q.Values.Del(key)
}
