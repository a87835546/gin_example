package controllers

import (
	"gin_example/param"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
)

type WatchedController struct {
	service *service.WatchedService
}

func NewWatchController() *WatchedController {
	return &WatchedController{
		service: service.NewWatchedService(),
	}
}
func (wc *WatchedController) GetListByUserId(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	list, err := wc.service.GetListByUserId(userId)
	if err != nil {
		RespError(ctx, utils.QueryDBErrorCode, err.Error())
	} else {
		RespOk(ctx, list)
	}
}
func (wc *WatchedController) AddWatch(ctx *gin.Context) {
	req := param.AddWatchReq{}
	err := ctx.ShouldBindJSON(&req)
	err = wc.service.AddWatch(&req)
	if err != nil {
		RespError(ctx, utils.InsertDBErrorCode, err.Error())
	} else {
		RespOk(ctx, nil)
	}
}
