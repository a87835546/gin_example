package param

type InsertReq struct {
	Title      string `json:"title" binding:"required"`
	Url        string `json:"url" binding:"required"`
	Desc       string `json:"desc" binding:"required"`
	Author     string `json:"author"`
	ThemeUrl   string `json:"theme_url" gorm:"colum:theme_url"`
	Types      string `json:"types"`
	Actor      string `json:"actor"`
	CategoryId string `json:"category_id"`
	Duration   int    `json:"duration"`
	Rate       string `json:"rate"`
	Years      int    `json:"years"`
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
	Position int    `json:"position,omitempty"UpdateMenuReq`
	Status   int    `json:"status"`
}
