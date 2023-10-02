package controllers

import (
	"gin_example/doreamon/utils"
	"gin_example/model"
	"gin_example/service"
	"github.com/gin-gonic/gin"
	"log"
)

type FavoriteController struct {
	sf *service.FavoriteService
}

func NewFavoriteController() *FavoriteController {
	fc := &FavoriteController{sf: service.NewFavoriteService()}
	return fc
}
func init() {

}
func (fc *FavoriteController) Add(ctx *gin.Context) {
	req := model.Favorite{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespError(ctx, utils.ParameterErrorCode, err.Error())
	} else {
		log.Printf("add fav req -->>%v\n", req)
		ok := fc.sf.Insert(&req)
		if ok {
			RespOk(ctx, nil)
		} else {
			RespError(ctx, utils.InsertDBErrorCode, nil)
		}
	}
}
func (fc *FavoriteController) Cancel(ctx *gin.Context) {
	req := model.Favorite{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		RespError(ctx, utils.ParameterErrorCode, err.Error())
	} else {
		ok := fc.sf.Cancel(&req)
		if ok {
			RespOk(ctx, nil)
		} else {
			RespError(ctx, utils.InsertDBErrorCode, nil)
		}
	}
}

func (fc *FavoriteController) QueryByUserId(ctx *gin.Context) {
	id, ok := ctx.GetQuery("user_id")
	if !ok {
		RespError(ctx, utils.ParameterErrorCode, nil)
	} else {
		list, err := fc.sf.QueryByUserId(id)
		if err == nil {
			RespOk(ctx, list)
		} else {
			RespError(ctx, utils.InsertDBErrorCode, nil)
		}
	}
}
func (fc *FavoriteController) QueryByVideoIdAndUserId(ctx *gin.Context) {
	req := make(map[string]string)
	err := ctx.ShouldBind(&req)
	if err != nil {
		RespError(ctx, utils.ParameterErrorCode, err.Error())
	} else {
		list, err := fc.sf.QueryByUserIdAndVideoId(req["user_id"], req["video_id"])
		if err == nil {
			RespOk(ctx, list)
		} else {
			RespError(ctx, utils.InsertDBErrorCode, nil)
		}
	}
}
