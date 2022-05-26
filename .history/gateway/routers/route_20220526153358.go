package routers

import (
	"net/http"
	"tiktok/pkg/logger"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RunServer(mode string) {
	// gin 设置成发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// TODOr.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")

	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	zap.L().Fatal(r.Run().Error())
}
