package response

import (
	"blog/pkg/enum"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type JsonResponse struct {
	Code     int         `json:"code"`    // 状态码
	Message  string      `json:"message"` // 消息内容
	Data     interface{} `json:"data"`    // 返回结构数据
	HttpCode int         `json:"-"`       // http状态码
}

// ToJson 响应json
func (resp *JsonResponse) ToJson(ctx *gin.Context) {
	code := http.StatusOK
	if resp.HttpCode != code {
		code = resp.HttpCode
	}
	ctx.JSON(code, resp)
}

// ToXml 响应Xml
func (resp *JsonResponse) ToXml(ctx *gin.Context) {
	code := http.StatusOK
	if resp.HttpCode != code {
		code = resp.HttpCode
	}
	ctx.XML(code, resp)
}

// ToString 响应String
func (resp *JsonResponse) ToString(ctx *gin.Context) {
	code := http.StatusOK
	if resp.HttpCode != code {
		code = resp.HttpCode
	}
	data, ok := resp.Data.(string)
	if !ok {
		panic("ToString---Data is not a string")
	}

	ctx.String(code, data)
}

// ToFile 响应File
func (resp *JsonResponse) ToFile(ctx *gin.Context) {
	path, ok := resp.Data.(string)
	if !ok {
		panic("ToFile---Data is not a string")
	}

	fileTmp, errByOpenFile := os.Open(path)
	defer func(fileTmp *os.File) {
		err := fileTmp.Close()
		if err != nil {
			panic("ToFile---File closure failed")
		}
	}(fileTmp)
	if errByOpenFile != nil {
		ctx.String(http.StatusNotFound, "file does not exist!")
	}

	ctx.File(path)
}

// SetHttpCode 设置HttpCode
func (resp *JsonResponse) SetHttpCode(httpCode int) *JsonResponse {
	resp.HttpCode = httpCode
	return resp
}

// Success 成功返回
func Success(message string, data ...interface{}) *JsonResponse {
	var r interface{}
	if len(data) > 0 {
		r = data[0]
	} else {
		r = struct{}{}
	}
	if message == "" {
		message = "Success"
	}
	return &JsonResponse{
		Code:    enum.HttpSuccess,
		Message: message,
		Data:    r,
	}
}

// Fail 失败返回
func Fail(code int, message string, data ...interface{}) *JsonResponse {
	var r interface{}
	if len(data) > 0 {
		r = data[0]
	} else {
		r = struct{}{}
	}
	return &JsonResponse{
		Code:    code,
		Message: message,
		Data:    r,
	}
}
