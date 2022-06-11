package controller

/*
func FeedHandler(c *gin.Context) {
	last_time := c.PostForm("last_time")
	token := c.PostForm("token")
	if last_time == "" {
		if token != "" {
			NowClaim, err := jwt.ParseToken(token)
			if err != nil {
				zap.L().Error("token is invalid", zap.Error(err))
				io.ResponseError(c, common.CodeTokenCreateErr)
				return
			}
		}
		return "video"
	} else {
		if token != "" {
			NowClaim, err := jwt.ParseToken(token)
			if err != nil {
				zap.L().Error("token is invalid", zap.Error(err))
				io.ResponseError(c, common.CodeTokenCreateErr)
				return
			}
		}
		return "timevideo"
	}

}*/
