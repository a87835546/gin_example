package controllers

import (
	"gin_example/param"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type BillboardController struct {
	vs service.BillboardService
	bs service.BannerService
	sc service.CategoryService
}

func NewBillboardController() *BillboardController {
	return &BillboardController{
		vs: service.BillboardService{},
		bs: service.BannerService{},
		sc: service.CategoryService{},
	}
}
func (mc *BillboardController) GetList(ctx *gin.Context) {
	list, err := mc.vs.GetList()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, 210, err.Error(), nil)
	}
}
func (mc *BillboardController) GetListByCategory(ctx *gin.Context) {
	title := ctx.Query("menu_id")
	categories, err := mc.sc.QueryByMenuId(title)
	if err != nil {
		return
	}
	temp := make([]*param.VideosType, 0)
	for i := 0; i < len(categories); i++ {
		c := categories[i]
		list, err := mc.vs.QueryByCategoryId(c.Id)
		if err == nil {
			vt := param.VideosType{
				Type: c.Title,
				List: list,
			}
			temp = append(temp, &vt)
		}
	}
	banner, err := mc.bs.QueryAllByMenuId(title)
	resp := param.VideosResp{
		Banner: banner,
		List:   temp,
	}
	if err == nil {
		RespOk(ctx, resp)
	} else {
		RespErrorWithMsg(ctx, 210, err.Error(), nil)
	}
}

func (mc *BillboardController) InsertBillboard(ctx *gin.Context) {
	req := param.InsertReq{}
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
				RespErrorWithMsg(ctx, 210, err.Error(), nil)
			}
		}
	}
}
func (mc *BillboardController) UpdateBillboard(ctx *gin.Context) {
	req := param.UpdateBillboardReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = mc.vs.Update(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, 210, err.Error(), nil)
		}
	}
}

func (mc *BillboardController) SearchBillboard(ctx *gin.Context) {
	req := param.SearchVideoReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, "获取参数异常", err.Error())
	} else {
		list, err := mc.vs.SearchByReq(req)
		if err == nil {
			RespOk(ctx, list)
		} else {
			RespErrorWithMsg(ctx, 210, err.Error(), nil)
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
