package controllers

import (
	param2 "gin_example/doreamon/param"
	"gin_example/doreamon/utils"
	"gin_example/model"
	"gin_example/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type BillboardController struct {
	vs *service.BillboardService
	bs service.BannerService
	sc service.CategoryService
	sh service.WatchedService
}

func NewBillboardController() *BillboardController {
	return &BillboardController{
		vs: service.NewBillboardService(),
		bs: service.BannerService{},
		sc: service.CategoryService{},
		sh: service.WatchedService{},
	}
}

func (mc *BillboardController) Query(ctx *gin.Context) {
	list, err := mc.vs.Query()
	if err != nil {
		RespErrorWithMsg(ctx, 201, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}

func (mc *BillboardController) GetList(ctx *gin.Context) {
	page := ctx.Query("page")
	num := ctx.Query("num")
	title := ctx.Query("menu_title")
	list, err := mc.vs.GetList(page, num, title)
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}

func (mc *BillboardController) Clicked(ctx *gin.Context) {
	uid := ctx.Query("user_id")
	vid := ctx.Query("video_id")
	uId, _ := strconv.Atoi(uid)
	vId, _ := strconv.Atoi(vid)

	req := model.WatchListModel{
		UserId:  uId,
		VideoId: int64(vId),
	}
	err := mc.sh.AddWatch(&req)
	if err == nil {
		RespOk(ctx, nil)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}

func (mc *BillboardController) GetListByCategory(ctx *gin.Context) {
	title := ctx.Query("menu_id")
	page := ctx.Query("page")
	num := ctx.Query("num")
	uid := ctx.Query("user_id")
	list, err := mc.vs.QueryByCategoryId(title, page, num)
	banner, err := mc.bs.QueryAllBannersByMenuIdWithUserId(title, uid)
	resp := param2.VideosResp{
		Banner: banner,
		List:   list,
	}
	if err == nil {
		RespOk(ctx, resp)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}
func (mc *BillboardController) GetVideoUrlsByVideoId(ctx *gin.Context) {
	vid := ctx.Query("video_id")
	id, err := mc.vs.QueryVideosUrlByVideoId(vid)
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
		return
	} else {
		RespOk(ctx, id)
	}
}
func (mc *BillboardController) InsertBillboard(ctx *gin.Context) {
	req := param2.InsertReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		log.Printf("billboard req--->> %#v", req)
		m, err := mc.vs.QueryByUrl(req.Url)
		if err == nil && m.Id > 0 {
			RespErrorWithMsg(ctx, utils.QueryDBErrorCode, "插入数据异常已经存在", m)
			return
		}
		log.Printf("billboard --->> %#v", m)
		if m.Id != 0 {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, "插入数据异常已经存在", m)
		} else {
			err = mc.vs.Insert(&req)
			if err == nil {
				RespOk(ctx, nil)
			} else {
				RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
			}
		}
	}
}
func (mc *BillboardController) UpdateBillboard(ctx *gin.Context) {
	req := param2.UpdateBillboardReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = mc.vs.Update(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
		}
	}
}
func (mc *BillboardController) QueryVideoByActor(ctx *gin.Context) {
	name := ctx.Query("name")
	list, err := mc.vs.QueryVideoByActor(name)
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}

func (mc *BillboardController) QuerySubVideoById(ctx *gin.Context) {
	id := ctx.Query("id")
	list, err := mc.vs.QuerySubVideoById(id)
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}
func (mc *BillboardController) SearchBillboard(ctx *gin.Context) {
	req := param2.SearchVideoReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, "获取参数异常", err.Error())
	} else {
		list, err := mc.vs.SearchByReq(req)
		if err == nil {
			RespOk(ctx, list)
		} else {
			RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
		}
	}
}

func (mc *BillboardController) Delete(ctx *gin.Context) {
	mp := make(map[string]int)
	err := ctx.ShouldBindJSON(&mp)
	if err != nil {
		return
	}
	id, _ := mp["id"]
	err = mc.vs.Delete(id)
	if err != nil {
		RespErrorWithMsg(ctx, utils.DeleteDBErrorCode, "删除数据异常", err.Error())
	} else {
		RespOk(ctx, nil)
	}
}

func (mc *BillboardController) VideoClick(ctx *gin.Context) {
	mp := make(map[string]any)
	err := ctx.BindJSON(&mp)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, "获取参数异常", err.Error())
		return
	} else {
		id, _ := mp["video_id"]
		uid, _ := mp["user_id"]
		err = mc.vs.InsertHistory(uid, id)
		if err != nil {
			RespErrorWithMsg(ctx, utils.DeleteDBErrorCode, "删除数据异常", err.Error())
		} else {
			RespOk(ctx, nil)
		}
	}
}
