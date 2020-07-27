package plover

import (
	"net/http"
	"strings"
)

type App struct {
	*Container
	Route    *Route
	NotFound NotFound
}

func (app *App) SetNotFound(NotFound NotFound) {
	app.NotFound = NotFound
}

func (app *App) dispatch(PATH string, methods []Controllers, middlewares []Middleware) bool {
	for _, control := range methods {
		if strings.EqualFold(control.PATH, PATH) {
			initInvoke(control.Handle, app)
			defer destructInvoke(control.Handle)
			action := control.Action
			var intercept bool
			if len(middlewares) > 0 && !control.WithoutMiddleware {
				for _, middleware := range middlewares {
					action, intercept = middleware(app, action)
					if intercept { // 是否拦截
						action()
						return true
					}
				}
			}
			action()
			return true
		}
	}
	return false
}

func (app *App) Crossdomain(allowed bool) {
	app.Set("Crossdomain", allowed)
}

func (app *App) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	app.Init(responseWriter, request)
	
	if Crossdomain, ok := app.Get("Crossdomain").(bool); Crossdomain && ok {
		app.Response.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
		app.Response.Header().Set("Access-Control-Allow-Headers", "authorization, language, cache-control, content-type, if-match, if-modified-since, if-none-match, if-unmodified-since, x-requested-with")
		app.Response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
		app.Response.Header().Set("Access-Control-Max-Age", "1728000")
		app.Response.Header().Set("Access-Control-Allow-Credentials", "true")
		app.Response.Header().Set("Access-Control-Expose-Headers", "token")
		if request.Method == http.MethodOptions {
			app.Response.JsonData(nil)
		}
	}
	
	methods := app.Route.GetMethod(strings.ToUpper(request.Method))
	middlewares := app.Route.GetMiddleware()
	
	PATH := app.Assets.AnalysisUri(request.URL.Path)
	
	if ok := app.dispatch(PATH, methods, middlewares); ok {
		return
	}
	// Any 路由匹配
	if ok := app.dispatch(PATH, app.Route.GetMethod("*"), middlewares); ok {
		return
	}
	
	// 分组调用
	for _, group := range app.Route.Groups {
		if group.HasBasic(PATH) {
			methods = group.GetMethod(strings.ToUpper(request.Method))
			middlewares = group.GetMiddleware()
			path := app.Assets.AnalysisUri(strings.TrimPrefix(PATH, group.basic))
			if ok := app.dispatch(path, methods, middlewares); ok {
				return
			}
			// Any 路由匹配
			if ok := app.dispatch(path, group.GetMethod("*"), middlewares); ok {
				return
			}
		}
	}
	
	if hasAssets := app.Assets.Handle(responseWriter, request); hasAssets {
		return
	}
	if app.NotFound != nil {
		app.NotFound(app)
		return
	}
	http.NotFound(responseWriter, request)
	return
}

func (app *App) Run(addr ...string) {
	var monitor string
	if len(addr) > 0 {
		monitor = addr[0]
	} else {
		var host string
		if strings.EqualFold(app.Config.GetString("app.host"), "") {
			host = "0.0.0.0"
		} else {
			host = app.Config.GetString("app.host")
		}
		var port string
		if strings.EqualFold(app.Config.GetString("app.port"), "") {
			port = "8888"
		} else {
			port = app.Config.GetString("app.port")
		}
		monitor = host + ":" + port
	}
	http.ListenAndServe(monitor, app)
}

func NewApp() *App {
	container := &Container{}
	container.Assets = new(Assets)
	container.Config = NewConfig()
	route := new(Route)
	app := &App{container, route, nil}
	return app
}
