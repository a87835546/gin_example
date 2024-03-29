package model

import (
	"time"
)

type User struct {
	Username   string    `json:"username" gorm:"colum:username"`
	Password   string    `json:"-"`
	Avatar     string    `json:"avatar"`
	Birthday   string    `json:"birthday"`
	Gender     int       `json:"gender"`
	Email      string    `json:"email"`
	Id         int       `json:"id"`
	DeviceType int       `json:"device_type"`
	Ip         string    `json:"-"`
	Token      string    `json:"token" gorm:"-"`
	Rule       string    `json:"-"`
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
	UserId     int   `json:"user_id"`
	VideoId    int64 `json:"video_id"`
	Id         int   `json:"id"`
	IsFavorite bool  `json:"is_favorite"`
	CreatedAt  int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type Admin struct {
	Username  string    `json:"username" gorm:"colum:username"`
	Password  string    `json:"-"`
	Id        int       `json:"id"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	Token     string    `json:"token" gorm:"-"`
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
	Id         int64                `json:"id" gorm:"id"`
	Url        string               `json:"url" gorm:"url"`
	Urls       []*VideoUrlListModel `json:"urls" gorm:"-"`
	Desc       string               `json:"desc"`
	Author     string               `json:"author"`
	Duration   int                  `json:"duration"`
	Rate       string               `json:"rate"`
	Years      string               `json:"years"`
	Title      string               `json:"title" gorm:"title"`
	Type       string               `json:"types" gorm:"column:types"`
	Actor      string               `json:"actor"`
	ThemeUrl   string               `json:"theme_url" gorm:"column:theme_url"`
	CategoryId int64                `json:"category_id" gorm:"category_id"`
	MenuTitle  string               `json:"menu_title" gorm:"menu_title"`
	CreatedAt  string               `json:"created_at" gorm:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at" gorm:"updated_at"`
}

type CategoryModel struct {
	Id        int64     `json:"id" gorm:"id"`
	CreatedAt int64     `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	Title     string    `json:"title" gorm:"title"`
	TitleEn   string    `json:"title_en" gorm:"title_en"`
	Desc      string    `json:"desc" gorm:"column:desc"`
	Index     int64     `json:"index" gorm:"column:index"`
	MenuId    int       `json:"menu_id" gorm:"menu_id"`
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
	Id        int    `json:"id" gorm:"id"`
	Title     string `json:"title" gorm:"title"`
	TitleEn   string `json:"title_en" gorm:"title_en"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	Author    string `json:"author" gorm:"column:author"`
}

type ActorModel struct {
	Id        int    `json:"id" gorm:"id"`
	Name      string `json:"name" gorm:"name"`
	NameEn    string `json:"name_en" gorm:"name_en"`
	AvatarUrl string `json:"avatar_url" gorm:"avatar_url"`
}

type WatchListModel struct {
	Id        int   `json:"id" gorm:"id"`
	VideoId   int64 `json:"video_id" gorm:"column:video_id"`
	UserId    int   `json:"user_id" gorm:"column:user_id"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type HistoryModel struct {
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
	Id        int   `json:"id"`
	UserId    int   `json:"user_id"`
	VideoId   int64 `json:"video_id"`
}

type BannerModel struct {
	CreatedAt     int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	Id            int    `json:"id"`
	VideoId       int64  `json:"video_id"`
	MenuId        int64  `json:"menu_id"`
	Title         string `json:"title" gorm:"title"`
	Desc          string `json:"desc" gorm:"desc"`
	Status        int    `json:"status" gorm:"status"`
	VideoUrl      string `json:"video_url" gorm:"video_url"`
	VideoThemeUrl string `json:"video_theme_url" gorm:"video_theme_url"`
	Operation     string `json:"operation" gorm:"operation"`
}

type BannerWithVideoModel struct {
	BannerModel
	Actor      string `json:"actor"`
	Years      int    `json:"years"`
	Title      string `json:"title" gorm:"title"`
	Type       string `json:"types" gorm:"column:types"`
	Rate       string `json:"rate"`
	CategoryId string `json:"category_id" gorm:"category_id"`
	MenuTitle  string `json:"menu_title" gorm:"menu_title"`
	IsFavorite bool   `json:"is_favorite"`
}

type VideoUrlListModel struct {
	Id        int    `json:"id"`
	VideoId   int64  `json:"video_id"`
	Url       string `json:"url"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
}

type Logs struct {
	Id        int    `json:"id"`
	Type      int    `json:"type"` // 0 管理后台的登录日志 1 用户的登录
	Desc      string `json:"desc"`
	Ip        string `json:"ip"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
}
