package controllers

import (
	"gin_example/doreamon/param"
	"gin_example/doreamon/utils"
	"gin_example/service"
	"github.com/gin-gonic/gin"
	"log"
)

type MenuController struct {
}

var ms = service.MenuService{}

func (mc *MenuController) GetMenus(ctx *gin.Context) {
	list, err := ms.GetMenus()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
	}
}

func (mc *MenuController) UpdateMenu(ctx *gin.Context) {
	req := param.UpdateMenuReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = ms.Update(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		}
	}
}

func (mc *MenuController) InsertMenus(ctx *gin.Context) {
	req := param.MenuInsertReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		m, err := ms.QueryByTitle(req.Title)
		log.Printf("menu --->> %#v", m)
		if m.Id != 0 {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, "插入数据异常已经存在", m)
		} else {
			err = ms.Insert(&req)
			if err == nil {
				RespOk(ctx, nil)
			} else {
				RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
			}
		}
	}

}
func (mc *MenuController) DeleteMenus(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	err := ms.Delete(id.(int))
	if err != nil {
		RespErrorWithMsg(ctx, utils.DeleteDBErrorCode, "删除数据异常", err.Error())
	} else {
		RespOk(ctx, nil)
	}
}
