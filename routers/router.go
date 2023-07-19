package routers

import (
	"gin_example/controllers"
	"gin_example/middleware"
	"github.com/gin-gonic/gin"
	_ "net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		usergroup := apiv1.Group("/user")
		user := controllers.UserCtl{}
		usergroup.POST("/add", user.AddUsers)
		usergroup.POST("/login", user.Login)
		usergroup.POST("/logout", user.Logout)
		usergroup.GET("/get", user.GetUsers)
		usergroup.GET("/test", user.FetchDataFromPron)
		usergroup.GET("/upload", user.Upload)
	}
	return r
}
