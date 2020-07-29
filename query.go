package plover

import (
	`encoding/json`
	"net/url"
	`strconv`
	"strings"
)

type Query struct {
	url.Values
}
// @title 是否有值
func (q Query) HasQuery() bool {
	return q.Values != nil
}
// @title 是否有值
// @return bool
func (q Query) Has(key string) bool {
	return !strings.EqualFold(q.Values.Get(key), "")
}

func (q Query) ToStruct(data interface{}, invoke ...func(key string, value interface{}) interface{}) (err error) {
	var jsonByte []byte
	var values = make(map[string]interface{})
	for key, value := range q.Values {
		if len(invoke) > 0 {
			Invoke := invoke[0]
			if len(value) > 1 {
				values[key] = Invoke(key, value)
			} else {
				values[key] = Invoke(key, value[0])
			}
		} else {
			if len(value) > 1 {
				values[key] = value
			} else {
				values[key] = value[0]
			}
		}
	}
	jsonByte, err = json.Marshal(values)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonByte, &data)
	if err != nil {
		return err
	}
	return nil
}

// @title 获取值
// @param key string
// @param defValue mixed
// @return mixed
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

// @title 获取int64数据
// @param key string
// @param defValue int64
// @return mixed
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

// @title 获取String数据
// @param key string
// @param defValue int64
// @return mixed
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
