package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
)

type MenuService struct {
}

func (ms *MenuService) GetMenus() (list []*models.MenuModel, err error) {
	rows, err := logic.Db.Table("menu").Rows()
	for rows.Next() {
		var l *models.MenuModel
		logic.Db.ScanRows(rows, &l)
		list = append(list, l)
	}
	return
}

func (ms *MenuService) Update(p *param.MenuInsertReq) error {
	err := logic.Db.Table("menu").Updates(p).Error
	return err
}
func (ms *MenuService) Insert(p *param.MenuInsertReq) error {
	err := logic.Db.Table("menu").Create(p).Error
	return err
}
func (ms *MenuService) Delete(id int) error {
	err := logic.Db.Table("menu").Delete(models.MenuModel{}).Where("id=?", id).Error
	return err
}
func (ms *MenuService) QueryByTitle(title string) (m *models.MenuModel, err error) {
	err = logic.Db.Debug().Table("menu").Find(&m).Where("title", title).Error
	return
}
