package bop

import (
	"gin-vue-admin/controller/api/bop"
	"github.com/gin-gonic/gin"
)

func InitChannelRouter(Router *gin.RouterGroup) {
	//LiveStreamRouter := Router.Group("livestream").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	LiveStreamRouter := Router.Group("channel")
	{
		LiveStreamRouter.POST("getChannelReportList", bop.GetChannelReportList) //获取直播列表
		LiveStreamRouter.POST("getAllChannelReports", bop.GetAllChannelReports) //获取直播列表

	}
}
