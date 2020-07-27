package plover

import (
	"strings"
)

type Group struct {
	basic       string
	middlewares []Middleware
	nodes       map[string][]Controllers
}

func (group *Group) GetMethod(method string) []Controllers {
	return group.nodes[method]
}

func (group *Group) GetMiddleware() []Middleware {
	return group.middlewares
}

func (group *Group) HasBasic(PATH string) bool {
	return strings.HasPrefix(PATH, group.basic)
}

func (group *Group) Use(middleware ...Middleware) {
	group.middlewares = middleware
}

func (group *Group) Add(method string, PATH string, handle Handle, action Action, withoutMiddleware []bool) *Group {
	if group.nodes == nil {
		group.nodes = make(map[string][]Controllers)
	}
	method = strings.ToUpper(method)
	PATH = strings.ToLower(strings.Trim(PATH, "/"))
	if strings.EqualFold(PATH, "") {
		PATH = "/"
	}
	var WithoutMiddleware bool
	if len(withoutMiddleware) > 0 {
		WithoutMiddleware = withoutMiddleware[0]
	} else {
		WithoutMiddleware = false
	}
	group.nodes[method] = append(group.nodes[method], Controllers{
		PATH:              PATH,
		Handle:            handle,
		Action:            action,
		WithoutMiddleware: WithoutMiddleware,
	})
	return group
}

func (group *Group) Get(PATH string, handle Handle, action Action, withoutMiddleware ...bool) *Group {
	return group.Add("GET", PATH, handle, action, withoutMiddleware)
}
func (group *Group) Post(PATH string, handle Handle, action Action, withoutMiddleware ...bool) *Group {
	return group.Add("POST", PATH, handle, action, withoutMiddleware)
}
func (group *Group) Put(PATH string, handle Handle, action Action, withoutMiddleware ...bool) *Group {
	return group.Add("PUT", PATH, handle, action, withoutMiddleware)
}
func (group *Group) Delete(PATH string, handle Handle, action Action, withoutMiddleware ...bool) *Group {
	return group.Add("DELETE", PATH, handle, action, withoutMiddleware)
}
func (group *Group) Any(PATH string, handle Handle, action Action, withoutMiddleware ...bool) *Group {
	return group.Add("*", PATH, handle, action, withoutMiddleware)
}
