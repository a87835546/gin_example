package param

import (
	"gin_example/model"
)

type VideosResp struct {
	Banner []*model.BannerWithVideoModel `json:"banner"`
	List   []*VideosType                 `json:"videos"`
}
type VideosType struct {
	Type   string             `json:"type"`
	TypeEn string             `json:"type_en"`
	List   []*model.Billboard `json:"list"`
}
type CategoryResp struct {
	Id          int64  `json:"id" gorm:"id"`
	CreatedAt   int64  `json:"created_at" gorm:"created_at"`
	Title       string `json:"title" gorm:"title"`
	TitleEn     string `json:"title_en" gorm:"title_en"`
	Desc        string `json:"desc" gorm:"column:desc"`
	Index       int64  `json:"index" gorm:"column:index"`
	MenuTitle   string `json:"menu_title" gorm:"menu_title"`
	MenuTitleEn string `json:"menu_title_en" gorm:"menu_title_en"`
}
