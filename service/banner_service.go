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
func (bs *BannerService) QueryAllByMenuId(id string) (list []*models.BannerModel, err error) {
	err = logic.Db.Debug().Table("banner").Where("menu_id=?", id).Find(&list).Error
	return
}
func (bs *BannerService) Insert(model *models.BannerModel) (err error) {
	err = logic.Db.Debug().Table("banner").Create(model).Error
	return
}
