package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
	"gorm.io/gorm"
	"time"
)

type CategoryService struct {
	Db *gorm.DB
}

func (ms *CategoryService) GetCategories() (list []*param.CategoryResp, err error) {
	err = logic.Db.Debug().Table("menu_category").Model(&param.CategoryResp{}).
		Select("menu_category.id,menu_category.created_at,menu_category.title,menu_category.title_en,menu.title as menu_title,menu.title_en as menu_title_en,menu_category.desc,menu_category.index").
		Joins("left join menu on menu_category.menu_id = menu.id and menu_category.status = 1").Find(&list).Error
	return
}

func (ms *CategoryService) GetCategoriesWithMenuByMenuId(id []string) (list []*param.CategoryResp, err error) {
	if len(id) == 0 {
		err = logic.Db.Table("menu_category").Model(&param.CategoryResp{}).
			Select("menu_category.id,menu_category.created_at,menu_category.title,menu_category.title_en,menu.title as menu_title,menu.title_en as menu_title_en,menu_category.desc,menu_category.index").
			Joins("left join menu on menu_category.menu_id = menu.id and menu_category.status = 1").Find(&list).Error
	} else {
		err = logic.Db.Table("menu_category").Model(&param.CategoryResp{}).
			Select("menu_category.id,menu_category.created_at,menu_category.title,menu_category.title_en,menu.title as menu_title,menu.title_en as menu_title_en,menu_category.desc,menu_category.index").
			Joins("left join menu on menu_category.menu_id = menu.id and menu_category.status = 1").Where("menu.id IN ?", id).Find(&list).Error
	}
	return
}
func (ms *CategoryService) GetAppCategories() (list []*models.AppCategoryModel, err error) {
	db := logic.Db.Table("menu_category").Where("super_title=?", "").Find(&list).Group("index")
	return list, db.Error
}

func (ms *CategoryService) EditAppCategories(model *models.AppCategoryModel) (err error) {
	err = logic.Db.Table("menu_category").Updates(model).Error
	if err != nil {
		err = logic.Db.Table("menu_category").Create(model).Error
	}
	return
}
func (ms *CategoryService) DeleteAppCategories(id int) (err error) {
	db := logic.Db.Table("menu_category").Where("id", id).Delete(models.AppCategoryModel{})
	return db.Error
}
func (ms *CategoryService) Update(p *models.CategoryModel) error {
	err := logic.Db.Table("menu_category").Updates(p).Error
	return err
}
func (ms *CategoryService) Insert(p *models.CategoryModel) error {
	p.CreatedAt = time.Now().UnixMilli()
	err := logic.Db.Table("menu_category").Create(p).Error
	return err
}
func (ms *CategoryService) Delete(id int) error {
	err := logic.Db.Table("menu_category").Where("id=?", id).Delete(models.CategoryModel{}).Error
	return err
}
func (ms *CategoryService) QueryByTitleWithId(title string, id int) (m *models.CategoryModel, err error) {
	err = logic.Db.Debug().Table("menu_category").Where("title=? and menu_id = ?", title, id).First(&m).Error
	return
}
func (ms *CategoryService) QueryByMenuId(id string) (m []*models.CategoryModel, err error) {
	err = logic.Db.Debug().Table("menu_category").Where("menu_id = ?", id).Find(&m).Error
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
	err = logic.Db.Table("video_type").Find(&list).Error
	return
}
