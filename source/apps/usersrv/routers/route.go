package routers

import (
	"net/http"
	"tiktok/base/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RunServer(mode string) {

	// gin 判断按照什么模式启动
	switch mode {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	}

	// 生成一个默认的路由引擎
	r := gin.New()

	// TODO:r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))

	// 注册定制的 gin 日志中间件，以及用于 gin 的 panic 后 recover 的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册接口路由
	registerRouter(r)

	// pprof.Register(r)

	// 访问不存在的路由就返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	// 路由启动失败，打印启动失败的日志（会中断）
	zap.L().Fatal(r.Run().Error())
}

// registerRouter 注册接口路由
func registerRouter(r *gin.Engine) {

	// 注册接口
	// r.POST("/usersrv/register/", controller.RegisterHandler)
	//r.POST("/usersrv/register/", controller.RegisterHandler)

	// 登录接口
	r.POST("/usersrv/login/")

}
