package service

import (
	"gin_example/logic"
	"gin_example/models"
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

func (ms *CategoryService) GetAppCategories() (list []*models.AppCategoryModel, err error) {
	db := logic.Db.Table("video_type").Find(&list).Group("index")
	return list, db.Error
}

func (ms *CategoryService) EditAppCategories(model *models.AppCategoryModel) (err error) {
	db := logic.Db.Table("video_type").Updates(&model)
	if db.Error != nil || db.RowsAffected == 0 {
		db = logic.Db.Table("video_type").Create(&model)
	}
	return db.Error
}
func (ms *CategoryService) DeleteAppCategories(id int) (err error) {
	db := logic.Db.Table("video_type").Where("id", id).Delete(models.AppCategoryModel{})
	return db.Error
}
func (ms *CategoryService) Update(p *models.CategoryModel) error {
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
