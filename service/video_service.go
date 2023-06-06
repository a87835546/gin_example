package service

import (
	"gin_example/logic"
	"gin_example/models"
	"log"
)

type VideoService struct {
}

func (vs *VideoService) Insert(videos []*models.VideoModel) error {
	_db := logic.Db.Table("tb_video").CreateInBatches(videos, len(videos))
	if _db.Error != nil {
		log.Printf("insert data err --->> %s\n", _db.Error.Error())
	}
	return _db.Error
}
func (vs *VideoService) InsertVideosInfo(video []*models.VideoInfo) error {
	_db := logic.Db.Table("tb_video_info").CreateInBatches(video, len(video))
	if _db.Error != nil {
		log.Printf("insert data err --->> %s\n", _db.Error.Error())
	}
	return _db.Error
}

func (vs *VideoService) UpdateNewUrl(url string) error {
	_db := logic.Db.Table("tb_video_info").UpdateColumn("new_url", url)
	if _db.Error != nil {
		log.Printf("insert data err --->> %s\n", _db.Error.Error())
	}
	return _db.Error
}
func (vs *VideoService) UploadVideos(video *models.VideoInfo) error {

	return nil
}
