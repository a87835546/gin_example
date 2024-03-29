package routers

import (
	"gin_example/controllers"
	"gin_example/doreamon"
	middleware2 "gin_example/doreamon/middleware"
	"gin_example/logic"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware2.Cors())
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
	r.Use(middleware2.ErrorHttp)
	//r.Use(doreamon.JWTAuth())
	//gin.DefaultWriter = io.MultiWriter(f)

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := logic.E.LoadPolicy()
	if err != nil {
		log.Printf("load policy err -->> %v", err)
	}
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/parser", controllers.Re)
		apiv1.GET("/parserOne", controllers.ParserOne)
		apiv1.GET("/parserOnePron", controllers.ParserOnePron)
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
			//billboardGroup.Use(middleware2.Auth(logic.E))
			bill := controllers.NewBillboardController()
			billboardGroup.GET("/list", bill.GetList)
			billboardGroup.GET("/click", bill.Clicked)
			billboardGroup.GET("/test", bill.Query)
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
			category := controllers.NewCategoryController()
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
			category := controllers.NewCategoryController()
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
			watchGroup.GET("/hot", wc.GetHotList)
			watchGroup.POST("/add", wc.AddWatch)
		}

		bannerGroup := apiv1.Group("/banner")
		{
			wc := controllers.NewBannerController()
			bannerGroup.GET("/list", wc.QueryAll)
			bannerGroup.GET("/queryByMenuId", wc.QueryByMenuId)
			bannerGroup.POST("/add", wc.Insert)
			bannerGroup.POST("/update", wc.Update)
		}

		ruleGroup := apiv1.Group("/rule")
		{
			rc := controllers.NewRuleController()
			ruleGroup.GET("/get", rc.GetPolicy)
			ruleGroup.GET("/getgroup", rc.GetGroupPolicy)
			ruleGroup.GET("/getgroupbyrule", rc.GetGroupPolicyByRule)
			ruleGroup.GET("/getGroupPolicy", rc.GetNamedGroupingPolicy)
			ruleGroup.GET("/delete", rc.DeleteGroupPolicy)
			ruleGroup.POST("/addPolicy", rc.AddPolicy)
			ruleGroup.POST("/addGroupPolicy", rc.AddGroupPolicy)

			ruleGroup.GET("/list", rc.GetRules)
			ruleGroup.GET("/getList", rc.GetRulesWithoutRoot)
			ruleGroup.POST("/insert", rc.InsertRule)
			ruleGroup.POST("/update", rc.UpdateRule)
			ruleGroup.GET("/users", rc.GetUsers)
			ruleGroup.POST("/user/add", rc.AddUsers)
			ruleGroup.POST("/user/update", rc.UpdateUser)

		}
		pGroup := apiv1.Group("/permission")
		{
			cc := controllers.NewPermissionController()
			pGroup.GET("/list", cc.Query)
			pGroup.POST("/add", cc.Add)
			pGroup.POST("/update", cc.Update)
		}
	}
	return r
}
