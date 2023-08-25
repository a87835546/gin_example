package service

import (
	"gin_example/logic"
	"gin_example/models"
	"log"
	"time"
)

type CategoryService struct {
}

func (ms *CategoryService) GetCategories() (list []*models.CategoryModel, err error) {
	rows, err := logic.Db.Debug().Table("video_category").Rows()
	for rows.Next() {
		var l *models.CategoryModel
		err = logic.Db.ScanRows(rows, &l)
		if err != nil {
			log.Println(err)
		}
		list = append(list, l)
	}
	return
}

func (ms *CategoryService) GetCategoriesBySuperId(id any) (list []*models.CategoryModel, err error) {
	rows, err := logic.Db.Table("video_category").Where("super_title=?", id).Rows()
	for rows.Next() {
		var l *models.CategoryModel
		err = logic.Db.ScanRows(rows, &l)
		list = append(list, l)
	}
	return
}

func (ms *CategoryService) GetAppCategories() (list []*models.AppCategoryModel, err error) {
	db := logic.Db.Table("video_category").Where("super_title=?", "").Find(&list).Group("index")
	return list, db.Error
}

func (ms *CategoryService) EditAppCategories(model *models.AppCategoryModel) (err error) {
	err = logic.Db.Table("video_category").Updates(model).Error
	if err != nil {
		err = logic.Db.Table("video_category").Create(model).Error
	}
	return
}
func (ms *CategoryService) DeleteAppCategories(id int) (err error) {
	db := logic.Db.Table("video_category").Where("id", id).Delete(models.AppCategoryModel{})
	return db.Error
}
func (ms *CategoryService) Update(p *models.CategoryModel) error {
	err := logic.Db.Table("video_category").Updates(p).Error
	return err
}
func (ms *CategoryService) Insert(p *models.CategoryModel) error {
	p.CreatedAt = time.Now().UnixMilli()
	err := logic.Db.Table("video_category").Create(p).Error
	return err
}
func (ms *CategoryService) Delete(id int) error {
	err := logic.Db.Table("video_category").Where("id=?", id).Delete(models.CategoryModel{}).Error
	return err
}
func (ms *CategoryService) QueryByTitle(title string) (m *models.CategoryModel, err error) {
	err = logic.Db.Debug().Table("video_category").Where("title=?", title).First(&m).Error
	return
}

func (ms *CategoryService) UpdateType(p *models.VideoTypeModel) error {
	err := logic.Db.Table("video_type").Updates(p).Error
	return err
}
func (ms *CategoryService) InsertType(p *models.VideoTypeModel) error {
	err := logic.Db.Table("video_type").Create(p).Error
	return err
}

func (ms *CategoryService) TypesBySuperId() (list []*models.VideoTypeModel, err error) {
	rows, err := logic.Db.Table("video_type").Rows()
	for rows.Next() {
		var l *models.VideoTypeModel
		logic.Db.ScanRows(rows, &l)
		list = append(list, l)
	}
	return
}
