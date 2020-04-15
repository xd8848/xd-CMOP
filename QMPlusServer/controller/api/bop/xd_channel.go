package bop

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/model/modelInterface"
	"gin-vue-admin/model/sysModel/bop"
	"github.com/gin-gonic/gin"
)

func GetChannelReportList(c *gin.Context) {
	// 此结构体仅本方法使用
	type searchParams struct {
		bop.XdChannelReport
		modelInterface.PageInfo
	}
	var sp searchParams
	_ = c.ShouldBindJSON(&sp)
	err, list, total := sp.XdChannelReport.GetInfoList(sp.PageInfo)
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"list":     list,
			"total":    total,
			"page":     sp.PageInfo.Page,
			"pageSize": sp.PageInfo.PageSize,
		})

	}
}
func GetAllChannelReports(c *gin.Context) {
	err, apis := new(bop.XdChannelReport).GetAllChannelReports()
	if err != nil {
		servers.ReportFormat(c, false, fmt.Sprintf("获取数据失败，%v", err), gin.H{})
	} else {
		servers.ReportFormat(c, true, "获取数据成功", gin.H{
			"apis": apis,
		})
	}
}
