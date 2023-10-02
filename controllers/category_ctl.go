package controllers

import (
	"encoding/json"
	"fmt"
	"gin_example/doreamon/utils"
	"gin_example/model"
	"gin_example/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CategoryController struct {
	db *service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		db: service.NewCategoryService(),
	}
}

func (mc *CategoryController) GetSubCategoriesWithMenu(ctx *gin.Context) {
	id := ctx.Query("id")
	ids := make([]string, 0)
	if len(id) > 0 {
		ids = strings.Split(id, ",")
	}
	list, err := mc.db.GetCategoriesWithMenuByMenuId(ids)
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
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
	ctx.Query("")
	list, err := mc.db.GetCategories()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}
func (mc *CategoryController) GetAppTabbarCategories(ctx *gin.Context) {
	list, err := mc.db.GetAppCategories()
	if err == nil || len(list) == 0 {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}
func (mc *CategoryController) ModifyAppTabbarCategories(ctx *gin.Context) {
	app := model.AppCategoryModel{}
	err := ctx.BindJSON(&app)
	err = mc.db.EditAppCategories(&app)
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
	err = mc.db.DeleteAppCategories(id)
	if err == nil {
		RespOk(ctx, nil)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}

func (mc *CategoryController) Update(ctx *gin.Context) {
	req := model.CategoryModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = mc.db.Update(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		}
	}
}

func (mc *CategoryController) InsertCategory(ctx *gin.Context) {
	req := model.CategoryModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		m, _ := mc.db.QueryByTitleWithId(req.Title, req.MenuId)
		log.Printf("menu --->> %#v", m)
		if m.Id != 0 {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, "插入数据异常已经存在", m)
		} else {
			err1 := mc.db.Insert(&req)
			if err1 == nil {
				RespOk(ctx, nil)
			} else {
				RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err1.Error(), nil)
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
	req := model.VideoTypeModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = mc.db.InsertType(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		}
	}
}

func (mc *CategoryController) UpdateType(ctx *gin.Context) {
	req := model.VideoTypeModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = mc.db.UpdateType(&req)
		if err == nil {
			RespOk(ctx, nil)
		} else {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		}
	}
}

func (mc *CategoryController) GetTypes(ctx *gin.Context) {
	list, err := mc.db.TypesBySuperId()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	}
}
