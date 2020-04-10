package bop

import (
	"gin-vue-admin/controller/api/bop"
	"github.com/gin-gonic/gin"
)

func InitAnchorRouter(Router *gin.RouterGroup) {
	//AnchorRouter := Router.Group("anchor").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	AnchorRouter := Router.Group("anchor")
	{
		AnchorRouter.POST("getAnchorList", bop.GetAnchorList)//获取直播管理列表
		AnchorRouter.POST("getAllAnchors", bop.GetAllAnchors)
	}
}
