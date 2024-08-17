package middleware

import (
	"blog/pkg/config"
	"blog/pkg/enum"
	"blog/pkg/logger"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recover 全局异常捕获
func Recover(cxt *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 将异常输出到日志
			var errorString = errorToString(r)
			logger.Logger.Error(errorString)
			if config.GetConf().Server.Mode == "debug" {
				response.Fail(enum.HttpFail, errorString).SetHttpCode(http.StatusInternalServerError).ToJson(cxt)
			} else {
				response.Fail(enum.HttpFail, "服务器错误，请联系开发者！").SetHttpCode(http.StatusInternalServerError).ToJson(cxt)
			}
			cxt.Abort()
		}
	}()
	cxt.Next()
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
