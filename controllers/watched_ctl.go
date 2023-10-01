package controllers

import (
	"gin_example/models"
	"gin_example/param"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
	"strconv"
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
	userId := ctx.Query("user_id")
	uid, err := strconv.Atoi(userId)
	list, err := wc.service.GetListByUserId(uid)
	if err != nil {
		RespError(ctx, utils.QueryDBErrorCode, err.Error())
	} else {
		RespOk(ctx, list)
	}
}

func (wc *WatchedController) GetHotList(ctx *gin.Context) {
	list, err := wc.service.GetHotList()
	if err != nil {
		RespError(ctx, utils.QueryDBErrorCode, err.Error())
	} else {
		RespOk(ctx, list)
	}
}
func (wc *WatchedController) AddWatch(ctx *gin.Context) {
	req := param.AddWatchReq{}
	err := ctx.ShouldBindJSON(&req)
	err = wc.service.AddWatch(&models.WatchListModel{})
	if err != nil {
		RespError(ctx, utils.InsertDBErrorCode, err.Error())
	} else {
		RespOk(ctx, nil)
	}
}
