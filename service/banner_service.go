package service

import (
	"gin_example/logic"
	"gin_example/models"
)

type BannerService struct {
}

func (bs *BannerService) QueryAll() (list []*models.BannerModel, err error) {
	err = logic.Db.Debug().Table("banner").Find(&list).Error
	return
}
func (bs *BannerService) QueryAllByMenuId(id string) (list []*models.BannerWithVideoModel, err error) {
	err = logic.Db.Debug().Table("banner").Select("banner.*,billboard.actor,billboard.years,billboard.types,billboard.rate,billboard.menu_title,billboard.category_id").Joins("LEFT JOIN billboard ON banner.video_id = billboard.id").Where("menu_id=?", id).Find(&list).Error
	return
}
func (bs *BannerService) Insert(model *models.BannerModel) (err error) {
	err = logic.Db.Debug().Table("banner").Create(model).Error
	return
}
func (bs *BannerService) Update(model *models.BannerModel) (err error) {
	err = logic.Db.Debug().Table("banner").Updates(model).Error
	return
}
