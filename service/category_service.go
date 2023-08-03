package service

import (
	"gin_example/logic"
	"gin_example/models"
	"gin_example/utils"
	"github.com/goccy/go-json"
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
	res, err := logic.Client.Get(utils.AppTabBarRedisKey).Result()
	if err != nil || len(res) == 0 {
		rows, err := logic.Db.Table("system_config").Where("key=?", "tabbar").Rows()
		if err != nil {
			for rows.Next() {
				var l *models.AppCategoryModel
				logic.Db.ScanRows(rows, &l)
				list = append(list, l)
			}
		}
		if len(list) > 0 {
			b, _ := json.Marshal(list)
			logic.Client.Set(utils.AppTabBarRedisKey, string(b), -1)
		}
	} else {
		err = json.Unmarshal([]byte(res), &list)
	}
	return
}

func (ms *CategoryService) EditAppCategories(val string) (err error) {
	_, err = logic.Client.Del(utils.AppTabBarRedisKey).Result()
	if err != nil {
		logic.Db.Table("system_config").Update("value", val).Where("key=?", "tabbar")
		mp := make(map[string]string, 0)
		json.Unmarshal([]byte(val), &mp)
		list := make([]*models.AppCategoryModel, 0)
		for _, value := range mp {
			m := models.AppCategoryModel{}
			json.Unmarshal([]byte(value), &m)
			list = append(list, &m)
		}
		if len(list) > 0 {
			b, _ := json.Marshal(list)
			logic.Client.Set(utils.AppTabBarRedisKey, string(b), -1)
		}
	}
	return
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
