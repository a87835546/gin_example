package param

type InsertReq struct {
	Title      string `json:"title" binding:"required"`
	Url        string `json:"url" binding:"required"`
	Desc       string `json:"desc" binding:"required"`
	Author     string `json:"author"`
	ThemeUrl   string `json:"theme_url"`
	CategoryId int    `json:"category_id"`
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
	Desc     string `json:"desc,omitempty" binding:"required"`
	Role     int    `json:"role,omitempty"`
	Position int    `json:"position,omitempty"`
	Status   int    `json:"status"`
}
