package controllers

import "github.com/gin-gonic/gin"

type BannerController struct {
}

func NewBannerController() *BannerController {
	return &BannerController{}
}

func (bc *BannerController) QueryByMenuId(ctx *gin.Context) {
	RespOk(ctx, nil)
}
