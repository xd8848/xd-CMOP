package bop

import (
	"gin-vue-admin/controller/api/bop"
	"github.com/gin-gonic/gin"
)

func InitLiveStreamRouter(Router *gin.RouterGroup) {
	//LiveStreamRouter := Router.Group("livestream").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	LiveStreamRouter := Router.Group("livestream")
	{
		LiveStreamRouter.POST("getLiveStreamList", bop.GetLiveStreamList) //获取直播列表
		//LiveStreamRouter.POST("getLiveStreamList", bop.GetLiveStreamList) //获取直播明细
		//LiveStreamRouter.POST("getLiveStreamList", bop.GetLiveStreamList) //获取打赏流水
		//LiveStreamRouter.POST("getLiveStreamList", bop.GetLiveStreamList) //获取打赏排行榜

	}
}
