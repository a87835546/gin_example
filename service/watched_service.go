package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
)

type WatchedService struct {
}

func NewWatchedService() *WatchedService {
	return &WatchedService{}
}
func (_ *WatchedService) GetListByUserId(id int) (list []*param.WatchListResp, err error) {
	rows, err := logic.Db.Debug().Table("watch_list").Where("user_id=?", id).Joins("left join billboard on billboard.id=watch_list.watch_id").Rows()
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

func (_ *WatchedService) AddWatch(req *param.AddWatchReq) (err error) {
	m := models.WatchListModel{}
	err = logic.Db.Debug().Table("watch_list").Where(req).Find(&m).Error
	if err != nil || m.Id <= 0 {
		err = logic.Db.Debug().Table("watch_list").Create(req).Error
	} else {
		err = logic.Db.Debug().Table("watch_list").Updates(req).Error
	}
	return
}
