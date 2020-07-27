package plover

import (
	"net/http"
	
	"github.com/unrolled/render"
)

type Container struct {
	responseWriter http.ResponseWriter
	request        *http.Request
	
	Request  Request
	Response Response
	
	Assets *Assets
	Config *Config
	
	bind map[string]interface{}
}

func (container *Container) Has(key string) (ok bool) {
	if _, ok := container.bind[key]; ok {
		return ok
	}
	return
}

func (container *Container) Add(key string, value interface{}) {
	if container.bind == nil {
		container.bind = make(map[string]interface{})
	}
	container.bind[key] = value
}

func (container *Container) Set(key string, value interface{}) {
	if container.bind == nil {
		container.bind = make(map[string]interface{})
	}
	container.bind[key] = value
}

func (container *Container) Get(key string) interface{} {
	if value, ok := container.bind[key]; ok {
		return value
	}
	return nil
}

// @title 获取原生数据响应
// @return request *http.Request
func (container *Container) ResponseWriter() (request *http.Request) {
	request = container.request
	return
}

func (container *Container) HttpRequest() (write http.ResponseWriter) {
	write = container.responseWriter
	return write
}

// 初始化容器
// @param http.ResponseWriter responseWriter
// @param *http.Request request
func (container *Container) Init(responseWriter http.ResponseWriter, request *http.Request) {
	container.request = request
	container.responseWriter = responseWriter
	container.Request = Request{request}
	var directory string
	if container.Config.GetString("template.directory") != "" {
		directory = container.Config.GetString("template.directory")
	} else {
		directory = "template"
	}
	var suffix []string
	if container.Config.GetStringMap("template.suffix") != nil {
		suffix = container.Config.GetStringSlice("template.suffix")
	} else {
		suffix = []string{".html", ".tmpl"}
	}
	container.Response = Response{responseWriter, 200, render.New(render.Options{
		Directory:  directory,
		Extensions: suffix,
	})}
}

// 实例化容器
// @param http.ResponseWriter responseWriter
// @param *http.Request request
// @return *Container
func NewContainer(responseWriter http.ResponseWriter, request *http.Request) *Container {
	container := new(Container)
	container.Init(responseWriter, request)
	return container
}
