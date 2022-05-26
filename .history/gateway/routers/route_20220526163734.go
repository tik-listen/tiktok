package routers

import (
	"net/http"
	"tiktok/pkg/logger"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RunServer(mode string) {
	// gin set release mode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// TODO:r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	
	r.LoadHTMLFiles("./templates/index.html")

	pprof.Register(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	zap.L().Fatal(r.Run().Error())
}
func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", )
	apiRouter.GET("/user/", )
	apiRouter.POST("/user/register/", )
	apiRouter.POST("/user/login/", )
	apiRouter.POST("/publish/action/", )
	apiRouter.GET("/publish/list/", )

	// extra apis - I
	apiRouter.POST("/favorite/action/", )
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
