package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/base/io"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 参数校验

	// 调用逻辑
	// VedioList :=
	// 返回响应
	c.JSON(http.StatusOK, io.FeedResponse{
		Response:  io.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
