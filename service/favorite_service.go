package service

import (
	"gin_example/logic"
	"gin_example/model"
	"gorm.io/gorm"
	"log"
)

type FavoriteService struct {
	db *gorm.DB
}

func NewFavoriteService() *FavoriteService {
	return &FavoriteService{
		db: logic.Db.Table("favorite").Debug(),
	}
}

func (fs *FavoriteService) Insert(favorite *model.Favorite) bool {
	fav := model.Favorite{}
	err := fs.db.Where("user_id=? and video_id=?", favorite.UserId, favorite.VideoId).Find(&fav).Error
	if err != nil {
		log.Println("query favorite db err", err.Error())
	}
	if err == nil && fav.Id != 0 {
		err = fs.db.Where("user_id=? and video_id=?", favorite.UserId, favorite.VideoId).Update("is_favorite", !fav.IsFavorite).Error
	} else {
		err = fs.db.Create(&favorite).Error
	}
	if err != nil {
		log.Println("insert favorite db err", err.Error())
	}
	return err == nil
}

func (fs *FavoriteService) Cancel(favorite *model.Favorite) bool {
	err := fs.db.Delete(favorite).Error
	if err != nil {
		log.Println("delete favorite db err", err.Error())
	}
	return err != nil
}
func (fs *FavoriteService) QueryByUserId(userId string) (list []*model.Favorite, err error) {
	rows, err := fs.db.Where("user_id=?", userId).Rows()
	if err != nil {
		log.Println("insert favorite db err", err.Error())
	}
	for rows.Next() {
		var l *model.Favorite
		err = logic.Db.ScanRows(rows, &l)
		list = append(list, l)
	}
	return
}
func (fs *FavoriteService) QueryByUserIdAndVideoId(userId, videoId string) (f *model.Favorite, err error) {
	err = fs.db.Where("user_id=? and video_id=?", userId, videoId).Find(&f).Error
	if err != nil {
		log.Println("insert favorite db err", err.Error())
	}
	return
}
