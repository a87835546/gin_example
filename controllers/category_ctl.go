package controllers

import (
	"gin_example/models"
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
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}

func (mc *CategoryController) GetAppTabbarCategories(ctx *gin.Context) {
	list, err := cs.GetAppCategories()
	if err == nil || len(list) == 0 {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}
func (mc *CategoryController) ModifyAppTabbarCategories(ctx *gin.Context) {
	app := models.AppCategoryModel{}
	err := ctx.BindJSON(&app)
	err = cs.EditAppCategories(&app)
	if err == nil {
		RespOk(ctx, nil)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}

func (mc *CategoryController) DeleteAppTabbarCategories(ctx *gin.Context) {
	mp := make(map[string]int, 0)
	err := ctx.BindJSON(&mp)
	id := mp["id"]
	err = cs.DeleteAppCategories(id)
	if err == nil {
		RespOk(ctx, nil)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}

func (mc *CategoryController) Update(ctx *gin.Context) {
	req := models.CategoryModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = cs.Update(&req)
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
