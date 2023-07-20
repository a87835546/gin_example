package controllers

import (
	"gin_example/param"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type BillboardController struct {
}

var bs = service.BillboardService{}

func (mc *BillboardController) GetList(ctx *gin.Context) {
	list, err := bs.GetList()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, 210, err.Error(), nil)
	}
}

func (mc *BillboardController) InsertBillboard(ctx *gin.Context) {
	req := param.InsertReq{}
	log.Printf("req --->>%#v", req)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = bs.Insert(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, 210, err.Error(), nil)
		}
	}
}
func (mc *BillboardController) UpdateBillboard(ctx *gin.Context) {
	req := param.InsertReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = bs.Insert(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, 210, err.Error(), nil)
		}
	}
}
