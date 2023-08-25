package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
)

type BillboardService struct {
}

func (bs *BillboardService) GetList() (list []*models.Billboard, err error) {
	logic.Db.Table("billboard").Find(&list)
	return
}

func (bs *BillboardService) Insert(billboard *param.InsertReq) (err error) {
	err = logic.Db.Table("billboard").Create(billboard).Error
	return
}

func (bs *BillboardService) QueryByUrl(url string) (bill *models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("url=?", url).First(&bill).Error
	return
}
func (bs *BillboardService) QueryByTitle(title string) (bill *models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("title=?", title).First(&bill).Error
	return
}
func (bs *BillboardService) Update(billboard *param.UpdateBillboardReq) (err error) {
	err = logic.Db.Debug().Table("billboard").Updates(&billboard).Where("id", billboard.Id).Error
	return
}
func (bs *BillboardService) Search(title string) (list []*models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("title", title).Find(&list).Error
	return
}
func (bs *BillboardService) Delete(i int) (err error) {
	err = logic.Db.Table("billboard").Where("id=?", i).Delete(models.Billboard{}).Error
	return err
}
func (bs *BillboardService) QueryByCategory(title string) (bill *models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("category_id=?", title).First(&bill).Error
	return
}
