package bop

import (
	"fmt"
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type XdEarningReport struct {
	gorm.Model
	StartDate        		string `json:"startDate"`
	EndDate 		 		string `json:"endDate"`
	AccountingDate 			string `json:"accountingDate"`
	BalanceCoin 		    string `json:"balanceCoin"`
	AnchorName 				string `json:"anchorName"`
	State 					string `json:"state"`
	Name      				string `json:"name"`

}



// 获取所有api信息
func (a *XdEarningReport) GetAllEarnings() (err error, apis []XdEarningReport) {
	err = qmsql.DEFAULTDB.Find(&apis).Error
	return
}
// 分页获取数据  需要分页实现这个接口即可
func (a *XdEarningReport) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {

	fmt.Println("================")
	fmt.Println(a.Name)
	fmt.Println(a.StartDate)
	fmt.Println(a.EndDate)
	fmt.Println("================")
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []XdEarningReport
		model := qmsql.DEFAULTDB.Model(info)
		if a.Name != "" {
			model = model.Where("anchor_name LIKE ?", "%"+a.Name+"%")
			db = db.Where("anchor_name LIKE ?", "%"+a.Name+"%")
		}
		if a.StartDate != "" {
			model = model.Where("accounting_date > str_to_date('"+a.StartDate+"', '%Y-%m-%d')")
			db = db.Where("accounting_date > str_to_date('"+a.StartDate+"', '%Y-%m-%d')")
		}
		if a.EndDate != "" {
			model = model.Where("accounting_date < str_to_date('"+a.EndDate+"', '%Y-%m-%d')")
			db = db.Where("accounting_date < str_to_date('"+a.EndDate+"', '%Y-%m-%d')")
		}
		err = model.Find(&apiList).Count(&total).Error

		return err, apiList, total
	}
}

