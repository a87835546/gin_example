package controllers

import (
	"gin_example/doreamon/utils"
	"gin_example/model"
	"gin_example/service"
	"github.com/gin-gonic/gin"
	"log"
)

type BannerController struct {
	bs *service.BannerService
	sb *service.BillboardService
}

func NewBannerController() *BannerController {
	return &BannerController{
		bs: service.NewBannerService(),
		sb: service.NewBillboardService(),
	}
}

func (bc *BannerController) QueryByMenuId(ctx *gin.Context) {
	id := ctx.Query("menu_id")
	list, err := bc.bs.QueryAllByMenuId(id)
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}
func (bc *BannerController) QueryAll(ctx *gin.Context) {
	list, err := bc.bs.QueryAll()
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}
func (bc *BannerController) Insert(ctx *gin.Context) {
	req := model.BannerModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		url, err := bc.sb.QueryVideoIdByUrl(req.VideoUrl)
		if err != nil || url == 0 {
			RespErrorWithMsg(ctx, utils.QueryDBErrorCode, "先把视频添加后再设置", nil)
		} else {
			req.VideoId = url
			err = bc.bs.Insert(&req)
			if err != nil {
				RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
			} else {
				RespOk(ctx, nil)
			}
		}
	}
}
func (bc *BannerController) Update(ctx *gin.Context) {
	req := model.BannerModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		log.Printf("update banner request %v\n", req.Id)
		err = bc.bs.Update(&req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}
