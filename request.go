package plover

import (
	`encoding/json`
	`encoding/xml`
	`errors`
	`io/ioutil`
	"net/http"
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

// @title 获取RawBody
// @return rawBody string
// @return error
func (r *Request) RawBody() (string, error) {
	rawBody, err := ioutil.ReadAll(r.Body)
	return string(rawBody), err
}

// @title 获取rawJson
// @param data interface
// @return interface, error
func (r *Request) RawJson(data interface{}) (interface{}, error) {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(rawBody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		return nil, errors.New("Content Type Not Application Json")
	}
}

// @title 获取rawXml
// @param data interface
// @return interface, error
func (r *Request) RawXml(data interface{}) (interface{}, error) {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "application/xml") {
		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		xml.Unmarshal(rawBody, &data)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		return nil, errors.New("Content Type Not Application Xml")
	}
}

func (r *Request) analysisForms() (data Query, err error) {
	if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		r.ParseMultipartForm(r.ContentLength)
		query := make(map[string][]string)
		for key, value := range r.MultipartForm.Value {
			if v, ok := query[key]; ok {
				query[key] = append(v, value...)
			} else {
				query[key] = value
			}
		}
		return Query{query}, nil
	} else if strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
		err = r.ParseForm()
		data = Query{r.Form}
		return
	} else {
		return Query{nil}, errors.New("Content-Type Cannot Be Resolved")
	}
}

func (r *Request) Get() Query {
	return Query{r.URL.Query()}
}

func (r *Request) Post() (data Query, err error) {
	if r.Method == http.MethodPost {
		return r.analysisForms()
	} else {
		return Query{nil}, errors.New("非法操作")
	}
}

func (r *Request) Put() (data Query, err error) {
	if r.Method == http.MethodPut {
		return r.analysisForms()
	} else {
		return Query{nil}, errors.New("非法操作")
	}
}

func (r *Request) Delete() (data Query, err error) {
	if r.Method == http.MethodDelete {
		return r.analysisForms()
	} else {
		return Query{nil}, errors.New("非法操作")
	}
}