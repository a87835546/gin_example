package controllers

import (
	"gin_example/models"
	"gin_example/param"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type CategoryController struct {
}

var cs = service.CategoryService{}

func (mc *CategoryController) GetCategories(ctx *gin.Context) {
	list, err := cs.GetCategories()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
	}
}

func (mc *CategoryController) UpdateMenu(ctx *gin.Context) {
	req := param.MenuInsertReq{}
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

func (mc *CategoryController) InsertCategory(ctx *gin.Context) {
	req := models.CategoryModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		m, err := cs.QueryByTitle(req.Title)
		log.Printf("menu --->> %#v", m)
		if m.Id != 0 {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, "插入数据异常已经存在", m)
		} else {
			err = cs.Insert(&req)
			if err == nil {
				RespOk(ctx, nil)
			} else {
				RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
			}
		}
	}

}
func (mc *CategoryController) DeleteMenus(ctx *gin.Context) {
	id, _ := ctx.Get("id")
	err := ms.Delete(id.(int))
	if err != nil {
		RespErrorWithMsg(ctx, utils.DeleteDBErrorCode, "删除数据异常", err.Error())
	} else {
		RespOk(ctx, nil)
	}
}
