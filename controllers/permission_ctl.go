package controllers

import (
	"gin_example/doreamon/utils"
	"gin_example/model"
	"gin_example/service"
	"github.com/gin-gonic/gin"
)

type PermissionController struct {
	ps service.PermissionService
}

func NewPermissionController() *PermissionController {
	return &PermissionController{
		ps: service.PermissionService{},
	}
}
func (pc *PermissionController) Add(ctx *gin.Context) {
	req := model.NewPermissionModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err := pc.ps.Add(&req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}

func (pc *PermissionController) Query(ctx *gin.Context) {
	list, err := pc.ps.QueryAll()
	if err != nil {
		RespErrorWithMsg(ctx, utils.InsertDBErrorCode, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}

func (pc *PermissionController) Update(ctx *gin.Context) {
	req := model.NewPermissionModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err := pc.ps.Update(&req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}
