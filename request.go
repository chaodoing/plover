package plover

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Request struct {
	*http.Request
}

// 是否是Post
// @return bool
func (r *Request) IsPost() bool {
	return strings.EqualFold(r.Method, http.MethodPost)
}

// 是否是Put
// @return bool
func (r *Request) IsPut() bool {
	return strings.EqualFold(r.Method, http.MethodPut)
}

// 是否是Delete
// @return bool
func (r *Request) IsDelete() bool {
	return strings.EqualFold(r.Method, http.MethodDelete)
}

// 是否是Get
// @return bool
func (r *Request) IsGet() bool {
	return strings.EqualFold(r.Method, http.MethodGet)
}

// 是否是Options请求
// @return bool
func (r *Request) IsOptions() bool {
	return strings.EqualFold(r.Method, http.MethodOptions)
}

// 是否是上传
// @return bool
func (r *Request) IsUploader() bool {
	if !strings.EqualFold(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		length, _ := strconv.ParseInt(r.Header.Get("Content-Length"), 10, 64)
		r.ParseMultipartForm(length)
		return len(r.MultipartForm.File) > 0
	} else {
		return false
	}
}

// 获取上传文件
// @return interface
func (r *Request) File() map[string][]*multipart.FileHeader {
	if r.IsUploader() {
		if !strings.EqualFold(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			length, _ := strconv.ParseInt(r.Header.Get("Content-Length"), 10, 64)
			r.ParseMultipartForm(length)
			return r.MultipartForm.File
		} else {
			return nil
		}
	} else {
		return nil
	}
}

// 获取Post数据
// @return map[string][]string
func (r *Request) Post() Query {
	if r.IsPost() {
		if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			r.ParseMultipartForm(r.ContentLength)
			return Query{r.MultipartForm.Value}
		} else {
			r.ParseForm()
			return Query{r.Form}
		}
	} else {
		return Query{nil}
	}
}

// 获取Get数据
// @return map[string][]string
func (r *Request) Query() Query {
	return Query{r.URL.Query()}
}

// 获取Put数据
// @return map[string][]string
func (r *Request) Put() Query {
	if r.IsPut() {
		return r.Post()
	} else {
		return Query{nil}
	}
}

// 获取Delete数据
// @return map[string][]string
func (r *Request) Delete() Query {
	if r.IsDelete() {
		s, err := ioutil.ReadAll(r.Body) // 把  body 内容读入字符串 s
		if err != nil {
			log.Fatal(err)
		}
		data, err := url.ParseQuery(string(s))
		if err != nil {
			log.Fatal(err)
		}
		return Query{data}
	} else {
		return Query{nil}
	}
}

// 获取RawBody
// @return string
func (r *Request) RawBody() string {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

// 获取 RawJson
// @return interface
func (r *Request) RawJson(data interface{}) interface{} {
	if strings.EqualFold(r.Header.Get("Content-Type"), "application/json") {
		body := r.RawBody()
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {
			log.Fatal(err)
		}
		return data
	} else {
		return nil
	}
}

// 获取 RawXml
// @return interface
func (r *Request) RawXml(data interface{}) interface{} {
	if strings.EqualFold(r.Header.Get("Content-Type"), "application/xml") {
		body := r.RawBody()
		err := xml.Unmarshal([]byte(body), &data)
		if err != nil {
			log.Fatal(err)
		}
		return data
	} else {
		return nil
	}
}
