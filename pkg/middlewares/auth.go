package middlewares

import (
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//authHeader := c.Request.Header.Get("Authorization")
		//if authHeader == "" {
		//	controller.ResponseError(c, controller.CodeNeedLogin)
		//	c.Abort()
		//	return
		//}
		//// 按空格分割
		//parts := strings.SplitN(authHeader, " ", 2)
		//if !(len(parts) == 2 && parts[0] == "Bearer") {
		//	controller.ResponseError(c, controller.CodeInvalidToken)
		//	c.Abort()
		//	return
		//}
		//// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		//mc, err := jwt.ParseToken(parts[1])
		//if err != nil {
		//	controller.ResponseError(c, controller.CodeInvalidToken)
		//	c.Abort()
		//	return
		//}
		//// 将当前请求的userID信息保存到请求的上下文c上
		//c.Set(controller.CtxUserIDKey, mc.UserID)
		//
		//c.Next()
	}
}
