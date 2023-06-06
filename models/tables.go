package models

type User struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type VideoInfo struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	Desc      string `json:"desc" db:"desc"`
	Url       string `json:"url" db:"url"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}
type VideoModel struct {
	Id         int    `json:"id" db:"id"`
	IsDownload bool   `json:"is_download" db:"is_download"`
	Url        string `json:"url" db:"url"`
	NewUrl     string `json:"new_url" db:"new_url"`
	CreatedAt  int64  `json:"created_at" db:"created_at"`
}
