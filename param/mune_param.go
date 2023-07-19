package param

type InsertReq struct {
	Title      string
	Url        string
	Desc       string
	CategoryId int `json:"category_id"`
}
