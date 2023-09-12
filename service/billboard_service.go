package service

import (
	"fmt"
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
	"gorm.io/gorm"
	"log"
	"strings"
	"sync"
)

type BillboardService struct {
}

func (bs *BillboardService) GetList() (list []*models.Billboard, err error) {
	logic.Db.Table("billboard").Order("id desc").Find(&list)
	return
}

func (bs *BillboardService) Insert(billboard *param.InsertReq) (err error) {
	urls := strings.Split(billboard.Url, " ")
	if len(urls) > 1 {
		titles := strings.Split(urls[0], "$")
		if len(titles) == 2 {
			billboard.Url = titles[1]
		}
	}
	tx := logic.Db.Begin()
	err = tx.Table("billboard").Create(billboard).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
	tx1 := logic.Db.Begin()
	if len(urls) > 0 {
		temp := make([]*models.VideoUrlListModel, 0)
		for u := 0; u < len(urls); u++ {
			if len(urls[u]) > 0 {
				titles := strings.Split(urls[u], "$")
				temp = append(temp, &models.VideoUrlListModel{Url: titles[1], Title: titles[0], VideoId: billboard.Id})
			}
		}
		err = tx1.Debug().Table("video_url").CreateInBatches(temp, len(temp)).Error
	}
	err = tx1.Commit().Error
	if err != nil {
		tx1.Rollback()
	}
	return
}
func (bs *BillboardService) InsertUrls(tx *gorm.DB, urls []string, vid int64) (err error) {
	temp := make([]*models.VideoUrlListModel, 0)
	for u := 0; u < len(urls); u++ {
		if len(urls[u]) > 0 {
			temp = append(temp, &models.VideoUrlListModel{VideoId: vid, Url: urls[u]})
		}
	}
	err = logic.Db.Debug().Table("video_url").CreateInBatches(temp, len(temp)).Error
	return
}
func (bs *BillboardService) QueryByUrl(url string) (bill *models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("url=?", url).First(&bill).Error
	return
}
func (bs *BillboardService) QueryVideoIdByUrl(url string) (id int64, err error) {
	err = logic.Db.Table("billboard").Select("id").Where("url=?", url).Scan(&id).Error
	return
}

func (bs *BillboardService) QueryByTitle(title string) (bill *models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("title=?", title).First(&bill).Error
	return
}
func (bs *BillboardService) Update(billboard *param.UpdateBillboardReq) (err error) {
	err = logic.Db.Debug().Table("billboard").Updates(&billboard).Where("id", billboard.Id).Error
	return
}
func (bs *BillboardService) Search(title string) (list []*models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("title", title).Find(&list).Error
	return
}
func (bs *BillboardService) SearchByReq(req param.SearchVideoReq) (list []*models.Billboard, err error) {
	err = logic.Db.Table("billboard").Where("title=? OR actor IN ? OR types=?", req.Name, req.Name, req.Name).Find(&list).Error
	return
}
func (bs *BillboardService) QueryVideoByActor(name string) (list []*models.Billboard, err error) {
	names := strings.Split(name, ",")
	log.Printf("names --->>> %s", names)
	if len(names) > 1 && len(names[0]) > 0 {
		name = names[0]
	}
	str := fmt.Sprintf("FIND_IN_SET('%s'"+
		",%s)", name, "actor")
	log.Printf("str--->>> %s", str)
	err = logic.Db.Debug().Table("billboard").Where(str).Find(&list).Limit(5).Error
	return
}

func (bs *BillboardService) QuerySubVideoById(name string) (list *models.Billboard, err error) {
	err = logic.Db.Debug().Table("billboard").Where("id=?", name).Find(&list).Error
	if fmt.Sprintf("%d", list.Id) == name {
		var temp []*models.VideoUrlListModel
		err = logic.Db.Debug().Table("video_url").Where("video_id = ?", name).Find(&temp).Error
		list.Urls = temp
	} else {
		return nil, nil
	}
	return
}

func (bs *BillboardService) Delete(i int) (err error) {
	err = logic.Db.Table("billboard").Where("id=?", i).Delete(models.Billboard{}).Error
	return err
}
func (bs *BillboardService) QueryByCategoryId(id any) (resp []*models.Billboard, err error) {
	videos := make([]*models.Billboard, 0)
	err = logic.Db.Debug().Table("billboard").Where("category_id = ?", id).Find(&videos).Error
	wg := sync.WaitGroup{}
	for i := 0; i < len(videos); i++ {
		wg.Add(1)
		go func(video *models.Billboard) {
			var temp []*models.VideoUrlListModel
			err = logic.Db.Debug().Table("video_url").Where("video_id = ?", video.Id).Find(&temp).Error
			video.Urls = temp
			wg.Done()
		}(videos[i])
	}
	wg.Wait()
	return videos, err
}
func (bs *BillboardService) InsertHistory(userId, videoId any) (err error) {
	mp := make(map[string]any)
	mp["user_id"] = userId
	mp["video_id"] = videoId
	err = logic.Db.Table("history").Create(mp).Error
	return
}
