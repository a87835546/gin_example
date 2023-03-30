package routers

import (
	"gin_example/controllers"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		user := controllers.UserCtl{}
		apiv1.POST("/add", user.AddUsers)
		apiv1.GET("/get", user.GetUsers)
	}

	return r
}
