package models

import "time"

type User struct {
	Username  string    `json:"username" gorm:"colum:username"`
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
	Type       string    `json:"types"`
	Actor      string    `json:"actor"`
	ThemeUrl   string    `json:"theme_url" gorm:"column:theme_url"`
	CategoryId int       `json:"category_id" gorm:"category_id"`
	CreatedAt  string    `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}

type CategoryModel struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt int64     `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	Title     string    `json:"title" gorm:"title"`
	Desc      string    `json:"desc" gorm:"column:desc"`
	Index     int64     `json:"index" gorm:"column:index"`
	SuperId   string    `json:"super_id" gorm:"column:super_id"`
}

type AppCategoryModel struct {
	Id    int    `json:"id" gorm:"id"`
	Index int    `json:"index" gorm:"index"`
	Title string `json:"title" gorm:"title"`
	Desc  string `json:"desc" gorm:"desc"`
}
type VideoTypeModel struct {
	Id     int    `json:"id" gorm:"id"`
	Title  string `json:"title" gorm:"title"`
	Author string `json:"author" gorm:"column:author"`
}

type ActorModel struct {
	Id     int    `json:"id" gorm:"id"`
	Name   string `json:"name" gorm:"name"`
	Avatar string `json:"avatar" gorm:"avatar"`
}
