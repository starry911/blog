package middleware

import (
	cache2 "blog/pkg/cache"
	"blog/pkg/enum"
	"blog/pkg/jwt"
	"blog/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// JwtAuth jwt认证中间件
func JwtAuth(key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")

		if token == "" {
			response.Fail(enum.InvalidArgument, "您还未登录呢！").ToJson(c)
			c.Abort()
			return
		}

		// 解析token
		userID, err := jwt.ParseToken(token)
		if err != nil {
			response.Fail(enum.TokenFailure, "登录已失效，请重新登录！").ToJson(c)
			c.Abort()
			return
		}

		// 单点登录检查，key：为存储在缓存中的键名
		oldToken := cache2.CH.RedisConn.Get(fmt.Sprintf("%s:%d", key, userID)).Val()
		if oldToken != token {
			response.Fail(enum.TokenFailure, "您的帐号已在别的地方登录，您被迫退出！若不是您本人操作，请注意帐号安全！").ToJson(c)
			c.Abort()
			return
		}

		// 自定义检查数据库中用户是否存在

		// 将用户ID保存到请求的上下文中
		c.Set("userId", userID)
		c.Next()
		// 后续的处理函数可以用过c.Get("userId")来获取当前请求的用户信息
	}
}
