package response

import "github.com/gin-gonic/gin"

// GinContextE 是对Gin.Context 的扩展
type GinContextE struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	TOKEN_ERROR = 2
	API_ERROR   = 1
	SUCCESS     = 0
)

func (g *GinContextE) Response(httpCode, errCode int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
}

func (g *GinContextE) Ok(code int) {
	g.Response(200, code, "操作成功", map[string]interface{}{})
}

func (g *GinContextE) OkWithMessage(code int, message string) {
	g.Response(200, code, message, map[string]interface{}{})
}

func (g *GinContextE) OkWithData(code int, data interface{}) {
	g.Response(200, code, "操作成功", data)
}

func (g *GinContextE) OkWithDetailed(code int, data interface{}, message string) {
	g.Response(200, code, message, data)
}

func (g *GinContextE) Fail(code int) {
	g.Response(200, code, "操作失败", map[string]interface{}{})
}

func (g *GinContextE) FailWithMessage(code int, message string) {
	g.Response(200, code, message, map[string]interface{}{})
}

func (g *GinContextE) FailWithData(code int, data interface{}) {
	g.Response(200, code, "操作失败", data)
}

func (g *GinContextE) FailWithDetailed(code int, data interface{}, message string) {
	g.Response(200, code, message, data)
}
