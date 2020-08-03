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
		this.Json(1, "Fail", err.Error())
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

func (this *Controller) Json(Code int, Message string, data interface{}) {
	this.App.Response.Json(Code, Message, data)
}
