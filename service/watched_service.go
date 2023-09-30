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
	err = logic.Db.Debug().Select("history.*,billboard.author,"+
		"billboard.duration,billboard.rate,billboard.years,billboard.title,billboard.actor,billboard.theme_url,billboard.types,billboard.url").
		Table("history").Where("history.user_id=?", id).
		Joins("left join billboard on billboard.id=history.video_id").
		Find(&list).Error
	return
}

func (_ *WatchedService) AddWatch(req *models.WatchListModel) (err error) {
	err = logic.Db.Debug().Table("history").Clauses(clause.OnConflict{UpdateAll: true}).Create(&req).Error
	return
}
