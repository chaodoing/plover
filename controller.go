package plover

import (
	_ "github.com/go-sql-driver/mysql"
	`github.com/jinzhu/gorm`
	`log`
)

type Controller struct {
	App      *App
	DB       *gorm.DB
	Response Response
	Request  Request
}

// @title 初始化函数
func (this *Controller) Init(app *App) {
	this.App = app
	this.DB = this.Db(app.Config)
	this.Response = app.Response
	this.Request = app.Request
}

// @title 析构释放变量
func (this *Controller) Destruct() {
	if this.DB != nil {
		this.DB.Close()
	}
	this.App = nil
}

// @title 获取 *gorm.DB
// @description 获取数据库操作对象
// @param config *bird.Config 配置类
// @return *gorm.DB
func (this *Controller) Db(config *Config) *gorm.DB {
	if this.DB == nil {
		var (
			dialect string = config.GetString("database.dialect")
			schema  string = config.GetString("database.schema")
		)
		db, err := gorm.Open(dialect, schema)
		if err != nil {
			log.Fatal(err)
		}
		db.LogMode(config.GetBool("database.log"))
		this.DB = db
	}
	return this.DB
}

// @title 获取当前分页参数
// @description 输出分页数据 ()
// return (page, limit, total, offset int) 当前页面, 每页条数, 初始化总计条数(0), 查询偏移条数
func (this *Controller) Pagex() (page, limit, total, offset int) {
	page  = this.Request.Query().GetInt("page", 1)
	limit = this.Request.Query().GetInt("limit", 15)
	total = 0
	if page > 0 {
		offset = (page - 1) * limit
	} else {
		offset = 0
		page = 1
	}
	return
}

// @title 修改数据
// @description 修改数据 (db, data)
// @param db gorm.DB 要查询的表名称
// @param data interface 模型
func (this *Controller) Change(db *gorm.DB, data interface{}) {
	updated := db.Updates(data).RowsAffected > 0
	if updated {
		this.Message("修改数据成功")
	} else {
		this.Message("修改数据失败", 1)
	}
}

// @title 添加数据
// @description 添加数据 (db, data)
// @param db gorm.DB 要查询的表名称
// @param data interface 模型
func (this *Controller) Add(db *gorm.DB, data interface{}) {
	created := db.Create(&data).RowsAffected > 0
	if created {
		this.Message("添加数据成功")
	} else {
		this.Message("添加数据失败", 1)
	}
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
