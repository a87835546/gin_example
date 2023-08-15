package controllers

import (
	"gin_example/models"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type ActorController struct {
}

var as service.ActorService

func (ac *ActorController) QueryAll(ctx *gin.Context) {
	list, err := as.QueryAll()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
	RespOk(ctx, nil)
}

func (ac *ActorController) Insert(ctx *gin.Context) {
	req := models.ActorModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		m, err := as.QueryByName(req.Name)
		log.Printf("actor --->> %#v", m)
		if m.Id != 0 {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, "插入数据异常已经存在", m)
		} else {
			err = as.Insert(&req)
			if err == nil {
				RespOk(ctx, nil)
			} else {
				RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
			}
		}
	}
}

func (ac *ActorController) Update(ctx *gin.Context) {
	req := models.ActorModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = as.Update(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		}
	}
}
