package plover

import (
	"github.com/unrolled/render"
	"net/http"
)

type Response struct {
	http.ResponseWriter
	statusCode int
	render     *render.Render
}

func (response *Response) Reset(responser Response) {
	response.render = responser.render
	response.statusCode = response.statusCode
	response.ResponseWriter = responser.ResponseWriter
}

func (response *Response) StatusCode(statusCode int) *Response {
	response.statusCode = statusCode
	return response
}

// 输出分页数据
// @param interface data
// @param int page
// @param int limit
// @param int total
func (response *Response) Page(data interface{}, page int, limit int, total int) {
	response.render.JSON(response, response.statusCode, map[string]interface{}{
		"code":    0,
		"message": "Success",
		"data":    data,
		"page":    page,
		"limit":   limit,
		"total":   total,
	})
	return
}

// @title 输出Json内容
// @param code int
// @param message int
// @param interface data
func (response *Response) Xml(code int, message string, data interface{}) {
	response.render.XML(response, response.statusCode, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	})
	return
}

// @title 输出Json内容
// @param code int
// @param message int
// @param interface data
func (response *Response) Json(code int, message string, data interface{}) {
	response.render.JSON(response, response.statusCode, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	})
	return
}

// @title 输出Json数据
// @param data interface
func (response *Response) JsonData(data interface{}) {
	response.render.JSON(response, response.statusCode, data)
	return
}

// 渲染模板
// @param string    template
// @param interface data
func (response *Response) View(template string, data interface{}) {
	response.render.HTML(response, response.statusCode, template, data)
	return
}

// 输出string数据
// @param string data
func (response *Response) String(data string) {
	response.render.Text(response, response.statusCode, data)
	return
}
