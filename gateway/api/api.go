package api

import (
	"fmt"
	"net/http"
	"time"

	"gw/library"
	"gw/pkg/ds"
	"gw/pkg/dy"
	"gw/pkg/fw"
	"gw/util"
	"tiktok/conf"

	"github.com/gin-gonic/gin"
)

// 入口函数
func Run(c *gin.Context) {
	t := time.NewTimer(conf.RequestTimeOut * time.Second)
	//设置 global
	var glb G
	glb.Rch = make(chan string)
	glb.Ech = make(chan error)

	go func(c *gin.Context, glb *G) {
		glb.RequestTime = util.GetTime()

		//设置请求访问的数据
		if err := glb.SetInfo(c); err != nil {
			glb.Ech <- err
			return
		}

		//容错
		decay := dy.Decay{
			Open:      glb.Md.Decay,
			DecayTime: glb.Md.DecayTime,
			Ctx:       c,
		}
		decayBody := decay.Start()
		if decayBody != "" {
			glb.Rch <- decayBody
			return
		}

		//获取要访问的url
		dns := ds.Dns{
			Ds:  glb.Md.Dns,
			Pth: glb.Md.To,
			Ctx: c,
		}
		dns.GetRestUrl()
		glb.To = dns.To
		glb.Query = dns.Query

		//流量检查
		flow := fw.Flow{
			Path: glb.To,
			Num:  glb.Md.Flow,
		}
		if err := flow.Check(); err != nil {
			glb.Ech <- err
			return
		}

		//发起请求
		hp := library.HttpRequest{
			Method:    glb.Md.Method,
			To:        glb.To,
			Query:     glb.Query,
			Out:       glb.Md.Timeout,
			CacheTime: glb.Md.CacheTime,
		}

		//发起请求
		body, err := hp.Http()
		if err != nil {
			glb.Ech <- err
			return
		}

		//写入上下文,目前用于容错
		c.Set("RequestBody", body)

		glb.Rch <- body
	}(c, &glb)

	select {
	case rch := <-glb.Rch:
		c.String(http.StatusOK, rch)
	case ech := <-glb.Ech:
		c.String(http.StatusInternalServerError, fmt.Sprintln(ech))
	case <-t.C:
		c.String(http.StatusNotFound, "request time out")
	}

	t.Stop()
}
