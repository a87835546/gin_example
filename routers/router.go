package routers

import (
	"gin_example/controllers"
	"gin_example/doreamon"
	"gin_example/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(doreamon.JWTAuth())

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		usergroup := apiv1.Group("/user")
		{
			user := controllers.UserCtl{}
			usergroup.POST("/add", user.AddUsers)
			usergroup.POST("/login", user.Login)
			usergroup.GET("/info", user.GetUser)
			usergroup.POST("/logout", user.Logout)
			usergroup.GET("/get", user.GetUsers)
			usergroup.GET("/test", user.FetchDataFromPron)
			usergroup.GET("/upload", user.Upload)
		}

		menug := apiv1.Group("/menu")
		{
			menu := controllers.MenuController{}
			menug.GET("/list", menu.GetMenus)
		}

		billboardGroup := apiv1.Group("/billboard")
		{
			bill := controllers.BillboardController{}
			billboardGroup.GET("/list", bill.GetList)
			billboardGroup.POST("/update", bill.UpdateBillboard)
			billboardGroup.POST("/insert", bill.InsertBillboard)
		}
	}
	return r
}
