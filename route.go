package plover

import (
	"strings"
)

type Route struct {
	Nodes       map[string][]Controllers
	Middlewares []Middleware
	Groups      []*Group
}

func (route *Route) Use(middleware ...Middleware) {
	route.Middlewares = middleware
}

func (route *Route) GetMethod(method string) []Controllers {
	return route.Nodes[method]
}

func (route *Route) GetMiddleware() []Middleware {
	return route.Middlewares
}

func (route *Route) Group(basic string) *Group {
	group := new(Group)
	basic = strings.ToLower(strings.Trim(basic, "/"))
	if strings.EqualFold(basic, "") {
		basic = "/"
	}
	group.basic = basic
	if route.Groups == nil {
		route.Groups = make([]*Group, 0)
	}
	route.Groups = append(route.Groups, group)
	return group
}

func (route *Route) Add(method string, PATH string, handle Handle, action Action, withoutMiddleware []bool) {
	if route.Nodes == nil {
		route.Nodes = make(map[string][]Controllers)
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
	route.Nodes[method] = append(route.Nodes[method], Controllers{
		PATH:              PATH,
		Action:            action,
		Handle:            handle,
		WithoutMiddleware: WithoutMiddleware,
	})
}

func (route *Route) Get(PATH string, handle Handle, action Action, withoutMiddleware ...bool) {
	route.Add("GET", PATH, handle, action, withoutMiddleware)
}
func (route *Route) Post(PATH string, handle Handle, action Action, withoutMiddleware ...bool) {
	route.Add("POST", PATH, handle, action, withoutMiddleware)
}
func (route *Route) Put(PATH string, handle Handle, action Action, withoutMiddleware ...bool) {
	route.Add("PUT", PATH, handle, action, withoutMiddleware)
}
func (route *Route) Delete(PATH string, handle Handle, action Action, withoutMiddleware ...bool) {
	route.Add("DELETE", PATH, handle, action, withoutMiddleware)
}
func (route *Route) Any(PATH string, handle Handle, action Action, withoutMiddleware ...bool) {
	route.Add("*", PATH, handle, action, withoutMiddleware)
}
