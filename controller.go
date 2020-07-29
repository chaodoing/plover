package plover

import (
	_ "github.com/go-sql-driver/mysql"
	`github.com/jinzhu/gorm`
	`log`
)

type Controller struct {
	db       *gorm.DB
	App      *App
	Response Response
	Request  Request
}

// @title 分页调用函数
// @param invoke func(page, limit, offset int) (data interface{}, total int, err error)
// @rerurn void
func (this *Controller) Page(invoke func(page, limit, offset int) (data interface{}, total int, err error)) {
	page := int(this.Request.Get().GetInt64("page", 1))
	limit := int(this.Request.Get().GetInt64("limit", 15))
	var offset int
	if page > 0 {
		offset = (page - 1) * limit
	} else {
		offset = 0
		page = 1
	}
	data, total, err := invoke(page, limit, offset)
	if err != nil {
		this.Fail(err.Error())
	} else {
		this.App.Response.Page(data, page, limit, total)
	}
}

// @title 初始化函数
func (this *Controller) Init(app *App) {
	this.App = app
	this.db = this.Db()
	this.Response = app.Response
	this.Request = app.Request
}

// @title 析构释放变量
func (this *Controller) Destruct() {
	if this.db != nil {
		this.db.Close()
		this.db = nil
	}
	this.App = nil
}

// @title 获取 *gorm.DB
// @description 获取数据库操作对象
// @param config *Config 配置类
// @return *gorm.DB
func (this *Controller) Db() *gorm.DB {
	var config *Config = this.App.Config
	if this.db == nil {
		var (
			dialect string = config.GetString("database.dialect")
			schema  string = config.GetString("database.schema")
		)
		db, err := gorm.Open(dialect, schema)
		if err != nil {
			log.Fatal(err)
		}
		db.LogMode(config.GetBool("database.log"))
		this.db = db
	}
	return this.db
}

// @title 输出成功消息
// @description 输出成功消息 (data, message, code)
// @param data interface 成功数据
// @param message string 成功消息
// @param code    int    成功代码
func (this *Controller) Success(data interface{}, opts ...interface{}) {
	var (
		Message string
		Code    int
	)
	if len(opts) == 1 {
		Message = opts[0].(string)
		Code = 0
	} else if len(opts) == 2 {
		Message = opts[0].(string)
		Code = opts[1].(int)
	} else {
		Message = "Ok"
		Code = 0
	}
	this.App.Response.Json(Code, Message, data)
}

// @title 输出内容
// @description 输出内容 (message, [code])
// @param message string 输出的消息
// @param code int 错误代码
func (this *Controller) Message(message string, code ...int) {
	var Code int
	if len(code) > 0 {
		Code = code[0]
	} else {
		Code = 0
	}
	this.App.Response.Json(Code, message, nil)
}

// @title 错误输出
// @description 错误消息输出 (message, code, data)
// @param message string 错误消息
// @param code int 错误代码
// @param data interface{}  错误数据
func (this *Controller) Fail(message string, opts ...interface{}) {
	var (
		Data interface{}
		Code int
	)
	if len(opts) == 1 {
		Code = opts[0].(int)
	} else if len(opts) == 2 {
		Code = opts[0].(int)
		Data = opts[1]
	} else {
		Code = 1
		Data = nil
	}
	this.App.Response.Json(Code, message, Data)
}
