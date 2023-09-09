package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
)

type MenuService struct {
}

func (ms *MenuService) GetMenus() (list []*models.MenuModel, err error) {
	err = logic.Db.Table("menu").Find(&list).Error
	return
}

func (ms *MenuService) Update(p *param.UpdateMenuReq) error {
	err := logic.Db.Table("menu").Updates(p).Where("id", p.Id).Error
	return err
}
func (ms *MenuService) Insert(p *param.MenuInsertReq) error {
	err := logic.Db.Table("menu").Create(p).Error
	return err
}
func (ms *MenuService) Delete(id int) error {
	err := logic.Db.Table("menu").Where("id=?", id).Update("status", 1).Error
	return err
}
func (ms *MenuService) QueryByTitle(title string) (m *models.MenuModel, err error) {
	err = logic.Db.Debug().Table("menu").Where("title=?", title).First(&m).Error
	return
}
