package controllers

import (
	"gin_example/models"
	"gin_example/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var us = service.UserService{}

type UserCtl struct {
}

func (uc *UserCtl) GetUsers(ctx *gin.Context) {
	ctx.JSON(200, us.GetUsers())
}
func (uc *UserCtl) AddUsers(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println("user 请求", user)
	ctx.JSON(200, us.InsertUser(&user))
}
