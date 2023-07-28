package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
)

type CategoryService struct {
}

func (ms *CategoryService) GetCategories() (list []*models.CategoryModel, err error) {
	rows, err := logic.Db.Table("category").Rows()
	for rows.Next() {
		var l *models.CategoryModel
		logic.Db.ScanRows(rows, &l)
		list = append(list, l)
	}
	return
}

func (ms *CategoryService) Update(p *param.MenuInsertReq) error {
	err := logic.Db.Table("category").Updates(p).Error
	return err
}
func (ms *CategoryService) Insert(p *models.CategoryModel) error {
	err := logic.Db.Table("category").Create(p).Error
	return err
}
func (ms *CategoryService) Delete(id int) error {
	err := logic.Db.Table("category").Where("id=?", id).Delete(models.CategoryModel{}).Error
	return err
}
func (ms *CategoryService) QueryByTitle(title string) (m *models.CategoryModel, err error) {
	err = logic.Db.Debug().Table("category").Where("title=?", title).First(&m).Error
	return
}
