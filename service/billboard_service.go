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
func (bs *BillboardService) Update(billboard *param.InsertReq) (err error) {
	err = logic.Db.Table("billboard").Updates(&billboard).Error
	return
}
