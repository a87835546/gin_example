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
		apiv1.Group("/user", func(context *gin.Context) {
			user := controllers.UserCtl{}
			apiv1.POST("/add", user.AddUsers)
			apiv1.POST("/login", user.Login)
			apiv1.POST("/logout", user.Logout)
			apiv1.GET("/get", user.GetUsers)
			apiv1.GET("/test", user.FetchDataFromPron)
			apiv1.GET("/upload", user.Upload)
		})
	}
	return r
}
