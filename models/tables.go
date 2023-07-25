package models

import "time"

type User struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Id        int       `json:"id"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
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

type MenuModel struct {
	Id        int       `json:"id" db:"id"`
	Desc      string    `json:"desc"`
	Title     string    `json:"title"`
	Role      int       `json:"role"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

type Billboard struct {
	Id         int64     `json:"id" gorm:"id"`
	Url        string    `json:"url" gorm:"url"`
	Desc       string    `json:"desc"`
	Author     string    `json:"author"`
	Title      string    `json:"title" gorm:"title"`
	ThemeUrl   string    `json:"theme_url" gorm:"theme"`
	CategoryId int       `json:"category_id" gorm:"category_id"`
	CreatedAt  string    `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}
