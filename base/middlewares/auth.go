package middlewares

import (
	"github.com/gin-gonic/gin"
	"tiktok/base/common"
	"tiktok/base/io"
	"tiktok/base/jwt"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		// 获取 token
		authToken := c.Query("token")
		if authToken == "" {
			io.ResponseError(c, common.CodeNeedLogin)
			c.Abort()
			return
		}

		// 解析 Token
		mc, err := jwt.ParseToken(authToken)
		if err != nil {
			io.ResponseError(c, common.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的 userID 和 username 信息保存到请求的上下文c上
		c.Set(common.KCtxUserIDKey, mc.UserID)
		c.Set(common.KCtxUserNameKey, mc.Username)

		c.Next()
	}
}
