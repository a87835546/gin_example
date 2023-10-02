package service

import (
	"gin_example/logic"
	"gin_example/model"
	"log"
)

type VideoService struct {
}

func (vs *VideoService) Insert(videos []*model.VideoModel) error {
	_db := logic.Db.Table("tb_video").CreateInBatches(videos, len(videos))
	if _db.Error != nil {
		log.Printf("insert data err --->> %s\n", _db.Error.Error())
	}
	return _db.Error
}
func (vs *VideoService) InsertVideosInfo(video []*model.VideoInfo) error {
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
func (vs *VideoService) UploadVideos(video *model.VideoInfo) error {

	return nil
}
