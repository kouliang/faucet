package route

import (
	"github.com/gin-gonic/gin"

	"faucet-app/business/controller"
	"faucet-app/setting"
)

func Init(mode string, cfg *setting.LogConfig) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(ginLogger(mode, cfg), gin.Recovery())
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "404",
		})
	})

	tap := r.Group("/api/v1")
	tap.POST("/claim_reward", controller.ClaimHandler)
	tap.POST("/check_reward", controller.CheckHandler)
	tap.POST("/shareUrl", controller.GetShareUrl)
	tap.POST("/twitterUploadCode", controller.TwitterUploadCode)
	tap.POST("/check_all", controller.CheckAll)

	// 设置静态文件路径
	r.Static("/faucet", "./public")
	r.GET("/verify/twitter", controller.Twitter)

	return r
}
