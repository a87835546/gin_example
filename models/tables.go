package models

import (
	"time"
)

type User struct {
	Username   string `json:"username" gorm:"colum:username"`
	Password   string `json:"-"`
	Avatar     string `json:"avatar"`
	Birthday   string `json:"birthday"`
	Gender     int    `json:"gender"`
	Email      string
	Id         int       `json:"id"`
	DeviceType int       `json:"device_type"`
	Ip         string    `json:"-"`
	Token      string    `json:"token" gorm:"-"`
	CreatedAt  time.Time `json:"-" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}

type AppUserRegisterReq struct {
	Username   string `json:"username" `
	Password   string `json:"password"`
	DeviceType int    `json:"device_type"`
	Birthday   string `json:"birthday"`
	Ip         string `json:"ip"`
}

type Favorite struct {
	UserId    int       `json:"user_id"`
	VideoId   int64     `json:"video_id"`
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

type Admin struct {
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
	TitleEn   string    `json:"title_en" gorm:"title_en"`
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
	Duration   int       `json:"duration"`
	Rate       string    `json:"rate"`
	Years      int       `json:"years"`
	Title      string    `json:"title" gorm:"title"`
	Type       string    `json:"types" gorm:"column:types"`
	Actor      string    `json:"actor"`
	ThemeUrl   string    `json:"theme_url" gorm:"column:theme_url"`
	CategoryId string    `json:"category_id" gorm:"category_id"`
	CreatedAt  string    `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
}

type CategoryModel struct {
	Id         int64     `json:"id" gorm:"id"`
	CreatedAt  int64     `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"updated_at"`
	Title      string    `json:"title" gorm:"title"`
	TitleEn    string    `json:"title_en" gorm:"title_en"`
	Desc       string    `json:"desc" gorm:"column:desc"`
	Index      int64     `json:"index" gorm:"column:index"`
	SuperTitle string    `json:"super_title" gorm:"column:super_title"`
}

type AppCategoryModel struct {
	Id         int    `json:"id" gorm:"id"`
	Index      int    `json:"index" gorm:"index"`
	Title      string `json:"title" gorm:"title"`
	TitleEn    string `json:"title_en" gorm:"title_en"`
	Desc       string `json:"desc" gorm:"desc"`
	SuperTitle string `json:"super_title" gorm:"column:super_title"`
}
type VideoTypeModel struct {
	Id     int    `json:"id" gorm:"id"`
	Title  string `json:"title" gorm:"title"`
	Author string `json:"author" gorm:"column:author"`
}

type ActorModel struct {
	Id        int    `json:"id" gorm:"id"`
	Name      string `json:"name" gorm:"name"`
	NameEn    string `json:"name_en" gorm:"name_en"`
	AvatarUrl string `json:"avatar_url" gorm:"avatar_url"`
}

type WatchListModel struct {
	Id        int       `json:"id" gorm:"id"`
	VideoId   int64     `json:"video_id" gorm:"column:video_id"`
	UserId    int       `json:"user_id" gorm:"column:user_id"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
}

type HistoryModel struct {
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
	Id        int
	UserId    int
	VideoId   int64
}
