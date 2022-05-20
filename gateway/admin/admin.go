package admin

import (
	"gw/pkg/admin"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	// 路由访问转发的接口
	r.POST("/req/add/api", admin.Add)

	// api 动态生成转发接口的表添加
	r.POST("/req/add/group", admin.AddGroup)
}
