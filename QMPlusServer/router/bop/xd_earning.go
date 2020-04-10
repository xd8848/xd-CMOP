package bop

import (
	"gin-vue-admin/controller/api/bop"
	"github.com/gin-gonic/gin"
)

func InitEarningRouter(Router *gin.RouterGroup) {
	//EarningRouter := Router.Group("earning").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	EarningRouter := Router.Group("earning")
	{
		EarningRouter.POST("getEarningList", bop.GetEarningList) //获取收益列表

		EarningRouter.POST("getAllEarnings", bop.GetAllEarnings) // 获取所有收益
	}
}
