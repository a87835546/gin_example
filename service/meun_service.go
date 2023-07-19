package service

import (
	"gin_example/logic"
	"gin_example/models"
)

type MenuService struct {
}

func (ms *MenuService) GetMenus() (list []*models.MenuModel, err error) {
	rows, err := logic.Db.Table("menu").Rows()
	for rows.Next() {
		var l *models.MenuModel
		logic.Db.ScanRows(rows, l)
		list = append(list, l)
	}
	return
}
