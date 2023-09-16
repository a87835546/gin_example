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
	//h := fmt.Sprintf("-%d", time.Now().Hour())
	//t := time.Now().Format("2006-01-02")
	//path := "log/" + t + h + ".log"
	//f, err := os.Create(path)
	//if err != nil {
	//
	//	log.Printf("err-->>%s", err.Error())
	//}
	//log.SetOutput(f)
	//r.Use(gin.LoggerWithWriter(f))
	r.Use(gin.Recovery())
	r.Use(middleware.ErrorHttp)
	//r.Use(doreamon.JWTAuth())
	//gin.DefaultWriter = io.MultiWriter(f)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/parser", controllers.Re)
		apiv1.GET("/parserOne", controllers.ParserOne)
		apiv1.GET("/batch", controllers.BatchInsert)
		apiv1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		admingroup := apiv1.Group("/admin")
		{
			user := controllers.NewUserCtl()
			admingroup.Use(doreamon.JWTAuth())
			admingroup.POST("/add", user.AddUsers)
			admingroup.POST("/login", user.Login)
			admingroup.GET("/info", user.GetUser)
			admingroup.POST("/logout", user.Logout)
			admingroup.GET("/get", user.GetUsers)
			admingroup.GET("/test", user.FetchDataFromPron)
			admingroup.GET("/upload", user.Upload)
		}

		usergroup := apiv1.Group("/user")
		{
			user := controllers.NewUserCtl()
			usergroup.Use(doreamon.JWTAuth())
			usergroup.POST("/login", user.AppUserLogin)
			usergroup.POST("/register", user.AppCreateUser)
			usergroup.GET("/info", user.AppGetUser)
		}

		menug := apiv1.Group("/menu")
		{
			menu := controllers.MenuController{}
			menug.GET("/list", menu.GetMenus)
			menug.POST("/update", menu.UpdateMenu)
			menug.POST("/insert", menu.InsertMenus)
			menug.POST("/delete", menu.DeleteMenus)
		}

		billboardGroup := apiv1.Group("/videos")
		{
			bill := controllers.NewBillboardController()
			billboardGroup.GET("/list", bill.GetList)
			billboardGroup.GET("/watch", bill.VideoClick)
			billboardGroup.POST("/update", bill.UpdateBillboard)
			billboardGroup.POST("/insert", bill.InsertBillboard)
			billboardGroup.POST("/delete", bill.Delete)
			billboardGroup.POST("/search", bill.SearchBillboard)
			billboardGroup.GET("/queryList", bill.GetListByCategory)
			billboardGroup.GET("/queryByActor", bill.QueryVideoByActor)
			billboardGroup.GET("/queryById", bill.QuerySubVideoById)
			billboardGroup.GET("/urls/videoId", bill.GetVideoUrlsByVideoId)
		}

		categoriesGroup := apiv1.Group("/category")
		{
			category := controllers.CategoryController{}
			categoriesGroup.GET("/list", category.GetCategories)
			categoriesGroup.GET("/queryList", category.GetSubCategoriesWithMenu)
			categoriesGroup.GET("/app", category.GetAppTabbarCategories)
			categoriesGroup.POST("/modify", category.ModifyAppTabbarCategories)
			categoriesGroup.POST("/update", category.Update)
			categoriesGroup.POST("/delete", category.DeleteAppTabbarCategories)
			categoriesGroup.POST("/insert", category.InsertCategory)
		}

		typeGroup := apiv1.Group("/type")
		{
			category := controllers.CategoryController{}
			typeGroup.GET("/list", category.GetTypes)
			typeGroup.POST("/modify", category.UpdateType)
			typeGroup.POST("/insert", category.InsertType)
		}

		actorGroup := apiv1.Group("/actor")
		{
			category := controllers.ActorController{}
			actorGroup.GET("/list", category.QueryAll)
			actorGroup.POST("/modify", category.Update)
			actorGroup.POST("/insert", category.Insert)
		}

		favoriteGroup := apiv1.Group("/favorite")
		{
			fav := controllers.NewFavoriteController()
			favoriteGroup.GET("/queryByUserId", fav.QueryByUserId)
			favoriteGroup.GET("/queryByUserIdAndVideoId", fav.QueryByUserId)
			favoriteGroup.POST("/cancel", fav.Cancel)
			favoriteGroup.POST("/insert", fav.Add)
		}
		watchGroup := apiv1.Group("/watch")
		{
			wc := controllers.NewWatchController()
			watchGroup.GET("/list", wc.GetListByUserId)
			watchGroup.POST("/add", wc.AddWatch)
		}

		bannerGroup := apiv1.Group("/banner")
		{
			wc := controllers.NewBannerController()
			bannerGroup.GET("/list", wc.QueryAll)
			bannerGroup.GET("/queryByMenuId", wc.QueryByMenuId)
			bannerGroup.POST("/add", wc.Insert)
			bannerGroup.POST("/update", wc.Insert)
		}
	}
	return r
}
