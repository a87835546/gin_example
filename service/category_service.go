package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
	"gorm.io/gorm"
	"time"
)

type CategoryService struct {
	Db     *gorm.DB
	TypeDb *gorm.DB
	MenuBd *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		Db:     logic.Db.Debug().Table("menu_category"),
		TypeDb: logic.Db.Debug().Table("video_type"),
		MenuBd: logic.Db.Debug().Table("menu"),
	}
}
func (ms *CategoryService) GetCategories() (list []*param.CategoryResp, err error) {
	err = logic.Db.Debug().Table("menu_category").Model(&param.CategoryResp{}).
		Select("menu_category.id,menu_category.created_at,menu_category.title,menu_category.title_en,menu.title as menu_title,menu.title_en as menu_title_en,menu_category.desc,menu_category.index").
		Joins("left join menu on menu_category.menu_id = menu.id and menu_category.status = 1").Find(&list).Error
	return
}

func (ms *CategoryService) GetCategoriesWithMenuByMenuId(id []string) (list []*param.CategoryResp, err error) {
	if len(id) == 0 {
		err = logic.Db.Debug().Model(&param.CategoryResp{}).Raw("select menu_category.id,menu_category.created_at,menu_category.title,menu_category.title_en,menu.title as menu_title,menu.title_en as menu_title_en,menu_category.desc,menu_category.index\nfrom menu_category\nleft join menu on menu_category.menu_id = menu.id and menu_category.status = 1").Find(&list).Error
	} else {
		err = logic.Db.Debug().Model(&param.CategoryResp{}).Raw("select menu_category.id,menu_category.created_at,menu_category.title,menu_category.title_en,menu.title as menu_title,menu.title_en as menu_title_en,menu_category.desc,menu_category.index\nfrom menu_category\nleft join menu on menu_category.menu_id = menu.id and menu_category.status = 1\nwhere menu_category.menu_id in ?", id).Find(&list).Error
	}
	return
}
func (ms *CategoryService) GetAppCategories() (list []*models.AppCategoryModel, err error) {
	db := ms.Db.Where("super_title=?", "").Find(&list).Group("index")
	return list, db.Error
}

func (ms *CategoryService) EditAppCategories(model *models.AppCategoryModel) (err error) {
	err = ms.Db.Updates(model).Error
	if err != nil {
		err = ms.Db.Create(model).Error
	}
	return
}
func (ms *CategoryService) DeleteAppCategories(id int) (err error) {
	db := ms.Db.Where("id", id).Delete(models.AppCategoryModel{})
	return db.Error
}
func (ms *CategoryService) Update(p *models.CategoryModel) error {
	err := ms.Db.Updates(p).Error
	return err
}
func (ms *CategoryService) Insert(p *models.CategoryModel) error {
	p.CreatedAt = time.Now().UnixMilli()
	res := logic.Db.Table("menu_category").Debug().Create(p)
	return res.Error
}
func (ms *CategoryService) Delete(id int) error {
	err := ms.Db.Where("id=?", id).Delete(models.CategoryModel{}).Error
	return err
}
func (ms *CategoryService) QueryByTitleWithId(title string, id int) (m *models.CategoryModel, err error) {
	err = ms.Db.Where("title=? and menu_id = ?", title, id).First(&m).Error
	return
}
func (ms *CategoryService) QueryByMenuId(id string) (m []*models.CategoryModel, err error) {
	err = ms.Db.Where("menu_id = ?", id).Find(&m).Error
	return
}
func (ms *CategoryService) UpdateType(p *models.VideoTypeModel) error {
	err := logic.Db.Table("video_type").Debug().Updates(p).Error
	return err
}
func (ms *CategoryService) InsertType(p *models.VideoTypeModel) error {
	err := ms.TypeDb.Create(p).Error
	return err
}

func (ms *CategoryService) TypesBySuperId() (list []*models.VideoTypeModel, err error) {
	err = ms.TypeDb.Find(&list).Error
	return
}
