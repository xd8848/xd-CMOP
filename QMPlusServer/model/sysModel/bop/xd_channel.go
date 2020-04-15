package bop

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type XdChannelReport struct {
	gorm.Model
	StartDate        		string `json:"startDate"`
	EndDate 		 		string `json:"endDate"`
	CreateAt 		 		string `json:"createdAt"`
	SourceId 			string `json:"sourceId"`
	SourceName 		    string `json:"sourceName"`
	Times 				string `json:"times"`


}


// 获取所有api信息
func (a *XdChannelReport) GetAllChannelReports() (err error, apis []XdChannelReport) {
	err = qmsql.DEFAULTDB.Find(&apis).Error
	return
}

// 分页获取数据  需要分页实现这个接口即可
func (a *XdChannelReport) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {

	fmt.Println("================")
	fmt.Println(a.SourceName)
	fmt.Println(a.StartDate)
	fmt.Println(a.EndDate)
	fmt.Println("================")
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []XdChannelReport
		model := qmsql.DEFAULTDB.Model(info)
		if a.SourceName != "" {
			model = model.Where("source_name LIKE ?", "%"+a.SourceName+"%")
			db = db.Where("source_name LIKE ?", "%"+a.SourceName+"%")
		}
		if a.StartDate != "" {
			model = model.Where("create_at > str_to_date('"+a.StartDate+"', '%Y-%m-%d')")
			db = db.Where("create_at > str_to_date('"+a.StartDate+"', '%Y-%m-%d')")
		}
		if a.EndDate != "" {
			model = model.Where("create_at < str_to_date('"+a.EndDate+"', '%Y-%m-%d')")
			db = db.Where("create_at < str_to_date('"+a.EndDate+"', '%Y-%m-%d')")
		}
		err = model.Find(&apiList).Count(&total).Error

		return err, apiList, total
	}
}

