package controllers

import (
	"fmt"
	"gin_example/doreamon"
	"gin_example/doreamon/utils"
	"gin_example/model"
	"gin_example/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
	"log"
	"sync"
	"time"
)

type UserLoginReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type AppUserLoginReq struct {
	UserLoginReq
	DeviceType int8 `json:"device_type"`
}

type UserCtl struct {
	mu sync.RWMutex
	us *service.UserService
	ls *service.LogsService
}

func NewUserCtl() *UserCtl {
	return &UserCtl{
		us: service.NewUserService(),
		ls: &service.LogsService{},
	}
}
func (uc *UserCtl) GetUsers(ctx *gin.Context) {
	ctx.JSON(200, uc.us.GetUsers())
}
func (uc *UserCtl) GetUserById(ctx *gin.Context) {
	if id, ok := ctx.Get("id"); ok {
		user, err := uc.us.QueryUserById(id.(string))
		if err != nil {
			RespError(ctx, 201, err.Error())
		} else {
			RespOk(ctx, user)
		}
	}
}
func (uc *UserCtl) GetUser(ctx *gin.Context) {
	if user, ok := ctx.Get("claims"); ok {
		u := user.(*doreamon.CustomClaims)
		log.Println("user-->>", user)
		name, err := uc.us.QueryUserByName(u.Name)
		if err != nil {
			RespErrorWithMsg(ctx, utils.UnknownErrorCode, "unknown error", nil)
		} else {
			RespOk(ctx, name)
		}
	} else {
		username := ctx.GetString("username")
		name, err := uc.us.QueryUserByName(username)
		if err != nil {
			RespErrorWithMsg(ctx, utils.UnknownErrorCode, "unknown error", nil)
		} else {
			RespOk(ctx, name)
		}
	}
}
func (uc *UserCtl) Login(ctx *gin.Context) {
	req := UserLoginReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	log.Println("req--->>>", req)
	user, err := uc.us.QueryUserByName(req.Username)
	if err != nil {
		RespErrorWithMsg(ctx, utils.InsertDBErrorCode, err.Error(), nil)
	} else if user.Password != req.Password {
		RespErrorWithMsg(ctx, utils.LoginPasswordErrorCode, "password is wrong", nil)
	} else {
		Ip := ctx.ClientIP()
		uc.ls.InsertLog(0, "管理账号登录:"+req.Username, Ip)
		generateToken(ctx, user)
	}
}

func (uc *UserCtl) AppUserLogin(ctx *gin.Context) {
	req := AppUserLoginReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Println("err--->>>", err)
	}
	log.Println("req--->>>", req)
	user, err := uc.us.AppQueryUserByName(req.Username)
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	} else if user.Password != req.Password {
		RespErrorWithMsg(ctx, utils.LoginPasswordErrorCode, "password is wrong", nil)
	} else {
		err := uc.us.AppUpdateIp(req.Username, ctx.ClientIP())
		if err != nil {
			return
		}
		Ip := ctx.ClientIP()
		uc.ls.InsertLog(0, "app用户登录:"+req.Username, Ip)
		generateAppUserToken(ctx, user)
	}
}

func (uc *UserCtl) AppCreateUser(ctx *gin.Context) {
	req := model.AppUserRegisterReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		return
	}
	req.Ip = ctx.ClientIP()
	log.Println("req--->>>", req)
	user, err := uc.us.AppQueryUserByName(req.Username)
	if err != nil || user.Id == 0 {
		user, err := uc.us.AppCreate(&req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, err.Error(), nil)
		} else {
			generateAppUserToken(ctx, user)
		}
	} else {
		RespErrorWithMsg(ctx, utils.UnknownErrorCode, "account was exists", nil)
	}
}

func (uc *UserCtl) AppGetUser(ctx *gin.Context) {
	if user, ok := ctx.Get("claims"); ok {
		log.Println("user-->>", user)
		RespOk(ctx, user)
	} else {
		RespErrorWithMsg(ctx, utils.UnknownErrorCode, "unknown error", nil)
	}
}

func (uc *UserCtl) Logout(ctx *gin.Context) {
	ctx.JSON(200, uc.us.GetUsers())
}
func (uc *UserCtl) AddUsers(ctx *gin.Context) {
	user := model.Admin{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		log.Println("user 请求", user)
		RespOk(ctx, uc.us.InsertUser(&user))
	}
}

func (uc *UserCtl) Upload(ctx *gin.Context) {
	go func() {
		s, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"), // 替换自己账户的region
			Credentials: credentials.NewStaticCredentials(

				"AKIA57G7AAAXD7QQWS6N",
				"0HFi7DN3GvtQBLY0fQQCvEpA3PDwqxW3s0kHlzSI",
				"",
			), // Sessiontoken是进程相关，应该是连接中可以返回 （可为空）
		})
		if err != nil {
			log.Println("aws  failed", err)
		}
		res1 := Read("/Users/yicen/GolandProjects/" + "Mofos - Melody Marks Is Up For Some Fun After Getting Poked A Few Times On Her Big Natural Boobs.mp4")
		fileName, err1 := service.UploadFileToS3(s, res1, fmt.Sprintf("%d.mp4", time.Now().UnixMicro()))
		if err1 != nil {
			log.Println("Upload failed ", err1)
		} else {
			log.Println("upload success ", fileName)
		}
		svs := service.VideoService{}
		svs.UpdateNewUrl(fileName)
	}()
	mp := make(map[string]any)
	mp["code"] = 200
	mp["res"] = "请求成功"
	ctx.JSON(200, mp)
}
func (uc *UserCtl) FetchDataFromPron(ctx *gin.Context) {

	c := colly.NewCollector()
	urls := make([]string, 0)
	urls1 := make([]*model.VideoModel, 0)
	// Find and visit all links
	c.OnHTML("div.wrapper div.container div.frontListingWrapper ul li.pcVideoListItem div.phimage a", func(e *colly.HTMLElement) {
		//factId, err := strconv.Atoi(e.Attr("id"))
		//if err != nil {
		//	log.Println("Could not get id")
		//}
		if len(e.Attr("href")) > 0 {
			e.Request.Visit(e.Attr("href"))
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		if r.URL.String() != "https://www.pornhub.com" {
			urls = append(urls, r.URL.String())
			urls1 = append(urls1, &model.VideoModel{
				Url:        r.URL.String(),
				IsDownload: false,
			})
		}
		//r.Headers.Set("User-Agent", RandomString())
	})

	c.Visit("https://www.pornhub.com")

	infos := make([]*model.VideoInfo, 0)
	wg := sync.WaitGroup{}
	for _, url := range urls1 {
		wg.Add(1)
		go func(m *model.VideoModel) {
			defer uc.mu.RUnlock()
			res, err := runCommand("lux  -i " + m.Url)
			if err != nil {
				log.Println("err -->> ", err.Error(), m.Url)
			} else {
				log.Println(res)
			}

			uc.mu.RLock()
			infos = append(infos, &model.VideoInfo{Url: m.Url, Title: res})
			wg.Done()
		}(url)
	}
	wg.Wait()
	svs := service.VideoService{}
	err := svs.Insert(urls1)
	svs.InsertVideosInfo(infos)
	if err != nil {
		ctx.JSON(201, err.Error())
	} else {
		ctx.JSON(200, nil)
	}
}
