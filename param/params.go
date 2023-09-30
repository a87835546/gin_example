package param

type InsertReq struct {
	Id         int64    `json:"id"`
	Title      string   `json:"title" binding:"required"`
	Url        string   `json:"url" binding:"required"`
	Desc       string   `json:"desc" binding:"required"`
	Author     string   `json:"author"`
	ThemeUrl   string   `json:"theme_url" gorm:"colum:theme_url"`
	Types      string   `json:"types"`
	Actor      string   `json:"actor"`
	CategoryId int      `json:"category_id"`
	MenuTitle  string   `json:"menu_title" binding:"required"`
	MenuId     int      `json:"menu_id" binding:"required"`
	Duration   int      `json:"duration"`
	Rate       string   `json:"rate"`
	Years      string   `json:"years"`
	Urls       []string `json:"urls" gorm:"-"`
}

type UpdateBillboardReq struct {
	Id int `json:"id"`
	InsertReq
}
type UpdateMenuReq struct {
	Id int `json:"id"`
	MenuInsertReq
}

type MenuInsertReq struct {
	Title    string `json:"title,omitempty" binding:"required"`
	TitleEn  string `json:"title_en,omitempty" binding:"required"`
	Desc     string `json:"desc,omitempty" binding:"required"`
	Role     int    `json:"role,omitempty"`
	Position int    `json:"position,omitempty"`
	Status   int    `json:"status"`
}
type WatchListResp struct {
	VideoId  int64  `json:"video_id" gorm:"column:video_id"`
	UserId   int    `json:"user_id" gorm:"column:user_id"`
	Author   string `json:"author"`
	Duration int    `json:"duration"`
	Rate     string `json:"rate"`
	Years    int    `json:"years"`
	Title    string `json:"title" gorm:"title"`
	Actor    string `json:"actor"`
	Type     string `json:"types" gorm:"column:types"`
	ThemeUrl string `json:"theme_url" gorm:"column:theme_url"`
	Url      string `json:"url" gorm:"column:url"`
}
type AddWatchReq struct {
	VideoId string `json:"video_id" gorm:"column:video_id"`
	UserId  string `json:"user_id" gorm:"column:user_id"`
}
type SearchVideoReq struct {
	Name  string `json:"name,omitempty"` //电影的名称
	Type  int    `json:"type"`           // 电影的类型
	Actor string `json:"actor"`          // 主演
}
