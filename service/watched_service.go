package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
	"gorm.io/gorm/clause"
)

type WatchedService struct {
}

func NewWatchedService() *WatchedService {
	return &WatchedService{}
}
func (_ *WatchedService) GetListByUserId(id int) (list []*param.WatchListResp, err error) {
	rows, err := logic.Db.Debug().Table("history").Where("user_id=?", id).Joins("left join billboard on billboard.id=history.watch_id").Rows()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		m := param.WatchListResp{}
		err := rows.Scan(&m)
		if err != nil {
			return nil, err
		}
		list = append(list, &m)
	}
	return
}

func (_ *WatchedService) AddWatch(req *models.WatchListModel) (err error) {
	err = logic.Db.Debug().Table("history").Clauses(clause.OnConflict{UpdateAll: true}).Create(&req).Error
	return
}
