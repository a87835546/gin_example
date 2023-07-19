package controllers

import (
	"gin_example/service"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
}

var ms = service.MenuService{}

func (mc *MenuController) GetMenus(ctx *gin.Context) {
	list, err := ms.GetMenus()
	if err == nil {
		RespOk(ctx, list)
	} else {
		RespErrorWithMsg(ctx, 210, err.Error(), nil)
	}
}
