package service

import (
	"fmt"
	"gin_example/logic"
	"gin_example/models"
	"gin_example/param"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
)

type BillboardService struct {
	db     *gorm.DB
	vUrlDb *gorm.DB
	noneDb *gorm.DB
}

func NewBillboardService() *BillboardService {
	return &BillboardService{
		db:     logic.Db.Debug().Table("billboard"),
		vUrlDb: logic.Db.Debug().Table("video_url"),
		noneDb: logic.Db.Debug(),
	}
}

func (bs *BillboardService) Query() (list []*models.Billboard, err error) {
	err = bs.db.Limit(10).Offset(0).Find(&list).Error
	return
}

func (bs *BillboardService) GetList(page, num string) (list []*models.Billboard, err error) {
	p, err := strconv.Atoi(page)
	n, err := strconv.Atoi(num)
	if p == 0 {
		p = 1
	}
	if n == 0 {
		n = 5
	}
	err = bs.db.Order("id desc").Limit(n).Offset((p - 1) * n).Find(&list).Error
	return
}

func (bs *BillboardService) Insert(billboard *param.InsertReq) (err error) {
	urls := make([]string, 0)
	if len(billboard.Urls) > 0 {
		urls = billboard.Urls
	} else {
		urls = strings.Split(billboard.Url, " ")
		if len(urls) > 1 {
			titles := strings.Split(urls[0], "$")
			if len(titles) == 2 {
				billboard.Url = titles[1]
			}
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
				if len(titles) > 1 {
					temp = append(temp, &models.VideoUrlListModel{Url: titles[1], Title: titles[0], VideoId: billboard.Id})
				}
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
func (bs *BillboardService) InsertUrls(urls []string, vid int64) (err error) {
	temp := make([]*models.VideoUrlListModel, 0)
	for u := 0; u < len(urls); u++ {
		if len(urls[u]) > 0 {
			temp = append(temp, &models.VideoUrlListModel{VideoId: vid, Url: urls[u]})
		}
	}
	err = bs.vUrlDb.CreateInBatches(temp, len(temp)).Error
	return
}
func (bs *BillboardService) QueryByUrl(url string) (bill *models.Billboard, err error) {
	err = bs.db.Where("url=?", url).First(&bill).Error
	return
}
func (bs *BillboardService) QueryVideoIdByUrl(url string) (id int64, err error) {
	err = bs.db.Select("id").Where("url=?", url).Scan(&id).Error
	return
}

func (bs *BillboardService) QueryByTitle(title string) (bill *models.Billboard, err error) {
	err = bs.db.Where("title=?", title).First(&bill).Error
	return
}
func (bs *BillboardService) Update(billboard *param.UpdateBillboardReq) (err error) {
	err = bs.db.Updates(&billboard).Where("id", billboard.Id).Error
	return
}
func (bs *BillboardService) Search(title string) (list []*models.Billboard, err error) {
	err = bs.db.Where("title=?", title).Find(&list).Error
	return
}
func (bs *BillboardService) SearchByReq(req param.SearchVideoReq) (list []*models.Billboard, err error) {
	err = bs.db.Where("title=? OR actor IN ? OR types=?", req.Name, req.Name, req.Name).Find(&list).Error
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
	err = bs.db.Where(str).Limit(5).Find(&list).Error
	return
}

func (bs *BillboardService) QuerySubVideoById(name string) (list *models.Billboard, err error) {
	err = bs.db.Where("id=?", name).Find(&list).Error
	if fmt.Sprintf("%d", list.Id) == name {
		var temp []*models.VideoUrlListModel
		err = bs.vUrlDb.Where("video_id = ?", name).Find(&temp).Error
		list.Urls = temp
	} else {
		return nil, nil
	}
	return
}

func (bs *BillboardService) Delete(i int) (err error) {
	err = bs.db.Where("id=?", i).Delete(models.Billboard{}).Error
	return err
}
func (bs *BillboardService) QueryByCategoryId(id any, page, num string) (resp []*param.VideosType, err error) {
	p, err := strconv.Atoi(page)
	n, err := strconv.Atoi(num)
	if p == 0 {
		p = 1
	}
	if n == 0 {
		n = 5
	}
	ids := make([]*models.CategoryModel, 0)
	err = bs.noneDb.Raw("SELECT * FROM menu_category WHERE menu_id = ? ", id).Scan(&ids).Error
	for i := 0; i < len(ids); i++ {
		videos := make([]*models.Billboard, 0)
		v := &param.VideosType{}
		err = bs.noneDb.Raw("SELECT * FROM billboard WHERE category_id = ? limit ?,?", ids[i].Id, (p-1)*n, n).Scan(&videos).Error
		v.Type = ids[i].Title
		v.TypeEn = ids[i].TitleEn
		v.List = videos
		resp = append(resp, v)
	}
	return
}

func (bs *BillboardService) QueryVideosUrlByVideoId(id any) (resp []*models.VideoUrlListModel, err error) {
	err = bs.vUrlDb.Where("video_id=?", id).Scan(&resp).Error
	return
}
func (bs *BillboardService) QueryVideosWithUrlsByCategoryId(id any, page, num string) (resp []*models.Billboard, err error) {
	p, err := strconv.Atoi(page)
	n, err := strconv.Atoi(num)
	if p == 0 {
		p = 1
	}
	if n == 0 {
		n = 20
	}
	err = logic.Db.Debug().Raw("SELECT * FROM billboard WHERE category_id IN (SELECT id FROM menu_category WHERE menu_id = ?  GROUP BY id)", id).Scan(&resp).Limit(n).Offset((p - 1) * n).Error
	ids := make([]int64, 0)
	for i := 0; i < len(resp); i++ {
		ids = append(ids, resp[i].Id)
	}
	urls := make([]models.VideoUrlListModel, 0)
	err = logic.Db.Debug().Raw("SELECT * FROM video_url WHERE video_id IN (SELECT id FROM billboard WHERE category_id IN (SELECT id FROM menu_category WHERE menu_id = ? GROUP BY id) GROUP BY id)", id).Scan(&urls).Error
	for i := 0; i < len(urls); i++ {
		for j := 0; j < len(resp); j++ {
			if resp[j].Id == urls[i].VideoId {
				resp[j].Urls = append(resp[j].Urls, &urls[i])
			}
		}
	}
	return
}
func (bs *BillboardService) InsertHistory(userId, videoId any) (err error) {
	mp := make(map[string]any)
	mp["user_id"] = userId
	mp["video_id"] = videoId
	err = logic.Db.Table("history").Create(mp).Error
	return
}
