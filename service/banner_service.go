package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gorm.io/gorm"
	"log"
)

type BannerService struct {
	db *gorm.DB
}

func NewBannerService() *BannerService {
	return &BannerService{
		db: logic.Db.Debug().Table("banner"),
	}
}
func (bs *BannerService) reset() {
	bs.db = logic.Db.Debug().Table("banner")
}
func (bs *BannerService) QueryAll() (list []*models.BannerModel, err error) {
	log.Printf("db --->>>> %v\n", bs.db)
	err = bs.db.Find(&list).Error
	return
}
func (bs *BannerService) QueryAllByMenuId(id string) (list []*models.BannerWithVideoModel, err error) {
	err = logic.Db.Debug().
		Raw("select banner.*,billboard.actor,billboard.years,billboard.types,billboard.rate,billboard.menu_title,billboard.category_id from banner LEFT JOIN billboard ON banner.video_id = billboard.id where banner.menu_id = ?", id).
		Find(&list).Error
	return
}
func (bs *BannerService) Insert(model *models.BannerModel) (err error) {
	bs.reset()
	err = bs.db.Create(model).Error
	return
}
func (bs *BannerService) Update(model *models.BannerModel) (err error) {
	err = bs.db.Updates(model).Error
	log.Printf("db --->>>> %v\n", bs.db)
	bs.reset()
	return
}
