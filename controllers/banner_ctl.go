package controllers

import (
	"gin_example/models"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
)

type BannerController struct {
	bs service.BannerService
}

func NewBannerController() *BannerController {
	return &BannerController{
		bs: service.BannerService{},
	}
}

func (bc *BannerController) QueryByMenuId(ctx *gin.Context) {
	list, err := bc.bs.QueryAll()
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
	req := models.BannerModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = bc.bs.Insert(&req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}
