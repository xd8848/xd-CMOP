package bop

import (
	"gin-vue-admin/controller/servers"
	"gin-vue-admin/init/qmsql"
	"gin-vue-admin/model/modelInterface"
	"github.com/jinzhu/gorm"
)

type XdAnchorReport struct {
	gorm.Model
	Audit       string `json:"audit"`
	Broadcast   string `json:"broadcast"`
	Name        string `json:"name"`
	LogoImg     string `json:"logoImg"`
	GapImg      string `json:"gapImg"`
	ZoneImg     string `json:"zoneImg"`
	Description string `json:"description"`
	State       string `json:"state"`
}

// 获取所有api信息
func (a *XdAnchorReport) GetAnchorList() (err error, apis []XdAnchorReport) {
	err = qmsql.DEFAULTDB.Find(&apis).Error
	return
}

// 分页获取数据  需要分页实现这个接口即可
func (a *XdAnchorReport) GetInfoList(info modelInterface.PageInfo) (err error, list interface{}, total int) {


	// 封装分页方法 调用即可 传入 当前的结构体和分页信息
	err, db, total := servers.PagingServer(a, info)
	if err != nil {
		return
	} else {
		var apiList []XdAnchorReport
		model := qmsql.DEFAULTDB.Model(info)
		if a.Name != "" {
			model = model.Where("anchor_name LIKE ?", "%"+a.Name+"%")
			db = db.Where("anchor_name LIKE ?", "%"+a.Name+"%")
		}
		if a.Broadcast != "" {
			switch {
			case a.Broadcast == "Disable":

				model = model.Where("broadcast = -1")
				db = db.Where("broadcast = -1")
			case a.Broadcast == "Pending":

				model = model.Where("broadcast = 0")
				db = db.Where("broadcast = 0")
			case a.Broadcast == "Online":

				model = model.Where("broadcast = 1")
				db = db.Where("broadcast = 1")
			}

		}

		if a.Audit != "" {
			switch {
			case a.Audit == "Failure":

				model = model.Where("audit = -1")
				db = db.Where("audit = 1")
			case a.Audit == "Success":

				model = model.Where("audit = 1")
				db = db.Where("audit = 1")

			}
		}
		err = model.Find(&apiList).Count(&total).Error

		return err, apiList, total
	}
}
