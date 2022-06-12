package routers

import (
	"net/http"
	"tiktok/base/logger"
	"tiktok/base/middlewares"
	"tiktok/gateway/controller"

	"github.com/gin-contrib/pprof"
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

	// 注册 pprof 监控
	pprof.Register(r)

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
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	{
		// 注册接口
		apiRouter.POST("/user/register/", controller.RegisterHandler)

		// 登录接口
		apiRouter.POST("/user/login/", controller.LoginHandler)

		// 视频流接口 with 鉴权
		apiRouter.GET("/feed/").Use(middlewares.JWTAuthMiddleware())

		// 获取用户信息接口
		apiRouter.GET("/user/", controller.UserInfo).Use(middlewares.JWTAuthMiddleware())

		// 发布相关路由组 with 鉴权
		publish := apiRouter.Group("/publish").Use(middlewares.JWTAuthMiddleware())
		{
			// 发布操作
			publish.POST("/action")

			// 查看发布记录操作
			publish.GET("/list/")

		}

		// 喜欢相关路由组 with 鉴权
		favorite := apiRouter.Group("/favorite")
		{
			// 点赞操作
			favorite.POST("/action/", controller.VideoLikeAction)

			// 查看点赞记录操作
			favorite.GET("/list/")
		}

		// 评论相关路由组 with 鉴权
		comment := apiRouter.Group("/comment").Use(middlewares.JWTAuthMiddleware())
		{
			comment.POST("/action/", controller.CommentHandler)
			comment.GET("/list/")
		}

		// 关系相关路由组 with 鉴权
		relation := apiRouter.Group("/relation")
		{
			relation.POST("/action/")
			relation.GET("/follower/list/")
			relation.GET("/follow/list/")
		}
	}

}
