package plover

type Action func()
type NotFound func(app *App)
type Handle interface {
	Init(app *App)
	Destruct()
}
// @title 调用初始化函数
// @param handle Handle 接口
// @param app *App 应用数据
func initInvoke (handle Handle, app *App) {
	handle.Init(app)
}

// @title 析构函数调用
// @param handle Handle 接口
func destructInvoke(handle Handle) {
	handle.Destruct()
}

type Controllers struct {
	WithoutMiddleware bool
	PATH              string
	Handle            Handle
	Action            Action
}

type Middleware func(app *App, next Action) (handle Action, intercept bool)
