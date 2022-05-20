package api

import (
	"fmt"

	"gw/library"
	"gw/pkg/api"
	"gw/pkg/middle"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	list, err := library.Group("group")
	if err != nil {
		panic(fmt.Sprintf("Api Route Was Wrong Err Was %s", err))
	}

	//动态加载路由,根据mongoDB中的path加载
	for _, v := range list {
		pth := v.Group
		r.Any(fmt.Sprintf("%s%s", pth, "*action"), api.Run, middle.Body())
	}
}
