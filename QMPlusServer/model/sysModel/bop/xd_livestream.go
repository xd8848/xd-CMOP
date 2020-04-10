package bop

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type XdLiveStreamReport struct {
	gorm.Model
	StartDate      			  string `json:"startDate"`
	EndDate 				  string `json:"endDate"`
	Name       				  string `json:"name"`
	AnchorName 		      	  string `json:"anchorName"`
	Rate 				      string `json:"rate"`
	RewardNum 			      int32  `json:"rewardNum"`
	Duration 				  string `json:"duration"`

}


// 直播统计
func (a *XdLiveStreamReport) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []XdLiveStreamReport
		model := qmsql.DEFAULTDB.Model(info)
		if a.Name != "" {
			model = model.Where("anchor_name LIKE ?", "%"+a.Name+"%")
			db = db.Where("anchor_name LIKE ?", "%"+a.Name+"%")
		}
		err = model.Find(&apiList).Count(&total).Error

		return err, apiList, total
	}
}

//直播明细
func (a *XdLiveStreamReport) GetLiveStreamInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []XdLiveStreamReport
		model := qmsql.DEFAULTDB.Model(info)
		if a.Name != "" {
			model = model.Where("anchor_name LIKE ?", "%"+a.Name+"%")
			db = db.Where("anchor_name LIKE ?", "%"+a.Name+"%")
		}
		err = model.Find(&apiList).Count(&total).Error
		return err, apiList, total
	}
}

//打赏流水
func (a *XdLiveStreamReport) GetRewardReceiptList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []XdLiveStreamReport
		model := qmsql.DEFAULTDB.Model(info)
		if a.Name != "" {
			model = model.Where("anchor_name LIKE ?", "%"+a.Name+"%")
			db = db.Where("anchor_name LIKE ?", "%"+a.Name+"%")
		}
		err = model.Find(&apiList).Count(&total).Error

		return err, apiList, total
	}
}

//打赏排行
func (a *XdLiveStreamReport) GetTopRewardList(info modelInterface.PageInfo) (err error, list interface{}, total int) {
	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []XdLiveStreamReport
		model := qmsql.DEFAULTDB.Model(info)
		if a.Name != "" {
			model = model.Where("anchor_name LIKE ?", "%"+a.Name+"%")
			db = db.Where("anchor_name LIKE ?", "%"+a.Name+"%")
		}
		err = model.Find(&apiList).Count(&total).Error

		return err, apiList, total
	}
}