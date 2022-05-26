package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tiktok/gateway/io"
)


// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 参数校验

	
	c.JSON(http.StatusOK, io.FeedResponse{
		Response:  io.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
