package controllers

import (
	"encoding/json"
	"fmt"
	"gin_example/models"
	"gin_example/service"
	"gin_example/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type CategoryController struct {
}

var cs = service.CategoryService{}

func (mc *CategoryController) GetSubCategories(ctx *gin.Context) {
	//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
	mp := make(map[string]any, 0)
	err := ctx.ShouldBind(&mp)
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	} else {
		id, ok := mp["id"]
		if !ok {
			id = 0
		}
		list, err := cs.GetCategories(id)
		if err == nil {
			RespOk(ctx, list)
		} else {
			RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
		}
	}
}
func GetJsonData(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(data))

	var jsonData map[string]any // map[string]interface{}
	data, _ = ioutil.ReadAll(c.Request.Body)
	if e := json.Unmarshal(data, &jsonData); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}
	c.JSON(http.StatusOK, jsonData)
}
func (mc *CategoryController) GetCategories(ctx *gin.Context) {
	list, err := cs.GetCategories(0)
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

func (mc *CategoryController) InsertType(ctx *gin.Context) {
	req := models.VideoTypeModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		m, err := cs.QueryByTitle(req.Title)
		log.Printf("menu --->> %#v", m)
		if m.Id != 0 {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, "插入数据异常已经存在", m)
		} else {
			err = cs.InsertType(&req)
			if err == nil {
				RespOk(ctx, nil)
			} else {
				RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
			}
		}
	}

}

func (mc *CategoryController) UpdateType(ctx *gin.Context) {
	req := models.VideoTypeModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = cs.UpdateType(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		}
	}
}

func (mc *CategoryController) GetTypes(ctx *gin.Context) {
	list, err := cs.TypesBySuperId()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}
