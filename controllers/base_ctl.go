package controllers

import (
	"bytes"
	"fmt"
	"gin_example/doreamon"
	"gin_example/logic"
	"gin_example/models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gocolly/colly/v2"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func runCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	arr := strings.Split(string(output), "\n")
	for _, s := range arr {
		log.Println(s)
	}
	//log.Println(arr[2])
	//arr1 := strings.Split(arr[2], "\n")

	return arr[2], err
}
func Read(filepath string) []byte {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("read file fail", err)
		return nil
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("read to fd fail", err)
		return nil
	}

	return fd
}

type Result struct {
	code    int
	message string
	data    any
}

func RespOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
	ctx.Next()
}
func RespError(ctx *gin.Context, code int, data interface{}) {
	RespErrorWithMsg(ctx, code, "fail", data)
}

func RespErrorWithMsg(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(200, gin.H{"code": code, "message": message, "data": data})
	ctx.Next()
}

// 生成令牌
func generateToken(c *gin.Context, user *models.Admin) {
	j := &doreamon.JWT{
		SigningKey: []byte("newtrekWang"),
	}

	claims := doreamon.CustomClaims{
		ID:   user.Id,
		Name: user.Username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 过期时间 一小时
			Issuer:    "newtrekWang",            //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	log.Println(token)
	key := fmt.Sprintf("user:%d:token", user.Id)
	logic.Client.Set(key, token, 3600*time.Second)
	user.Token = token
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功！",
		"data":    user,
	})
	return
}
func generateAppUserToken(c *gin.Context, user *models.User) {
	j := &doreamon.JWT{
		SigningKey: []byte("newtrekWang"),
	}

	claims := doreamon.CustomClaims{
		ID:   user.Id,
		Name: user.Username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + 3600, // 过期时间 一小时
			Issuer:    "newtrekWang",            //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	log.Println(token)
	key := fmt.Sprintf("user:%d:token", user.Id)
	logic.Client.Set(key, token, 3600*time.Second)
	user.Token = token
	RespOk(c, user)
	return
}

func SystemConfig(ctx *gin.Context) {
	RespOk(ctx, nil)
}

func BatchInsert(ctx *gin.Context) {
	uid := ctx.Query("uid")
	idStr := ctx.Query("category_id")
	title := ctx.Query("menu_title")
	start := ctx.Query("start")
	end := ctx.Query("end")

	mid := ctx.Query("menu_id")
	mId, _ := strconv.Atoi(mid)

	id, _ := strconv.Atoi(idStr)
	s, _ := strconv.Atoi(start)
	e, _ := strconv.Atoi(end)

	for i := s; i < e; i++ {
		url := fmt.Sprintf("https://bfzy.tv/index.php/vod/type/id/%s/page/%d.html", uid, i)
		urls := make([]string, 0)
		id := id
		title := title
		c := colly.NewCollector(
			//colly.Async(true),
			colly.MaxDepth(2),
		)
		c.OnResponse(func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
		})
		c.OnHTML(".videoName", func(e *colly.HTMLElement) {
			fmt.Printf("videoName-->>>%s\n", e.Attr("href"))
			url := e.Attr("href")
			if len(url) > 0 {
				urls = append(urls, e.Attr("href"))
			}
		})
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err.Error())
		})

		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("User-Agent", RandomString())
			fmt.Println("Visiting", r.URL.String())
		})

		c.Visit(url)

		for i := 0; i < len(urls); i++ {
			url1 := fmt.Sprintf("https://bfzy.tv%s", urls[i])
			parserOne(url1, title, id, mId)
		}
	}

}

func Re(ctx *gin.Context) {
	urls := make([]string, 0)
	url := ctx.Query("url")
	id := ctx.Query("category_id")
	title := ctx.Query("menu_title")
	categoryId, _ := strconv.Atoi(id)

	mid := ctx.Query("menu_id")
	mId, _ := strconv.Atoi(mid)

	c := colly.NewCollector(
		//colly.Async(true),
		colly.MaxDepth(2),
	)
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnHTML(".videoName", func(e *colly.HTMLElement) {
		fmt.Printf("videoName-->>>%s\n", e.Attr("href"))
		url := e.Attr("href")
		if len(url) > 0 {
			urls = append(urls, e.Attr("href"))
		}
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err.Error())
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	for i := 0; i < len(urls); i++ {
		url1 := fmt.Sprintf("https://bfzy.tv%s", urls[i])
		parserOne(url1, title, categoryId, mId)
	}
}

type Video struct {
	ThemeUrl   string   `json:"theme_url,omitempty"`
	Urls       string   `json:"url,omitempty"`
	URLs       []string `json:"urls,omitempty"`
	Title      string   `json:"title,omitempty"`
	Desc       string   `json:"desc,omitempty"`
	Actor      string   `json:"actor,omitempty"`
	Rate       string   `json:"rate,omitempty"`
	Year       string   `json:"years,omitempty"`
	Types      string   `json:"types,omitempty"`
	MenuTitle  string   `json:"menu_title"`
	MenuId     int      `json:"menu_id"`
	Author     string   `json:"author"`
	CategoryId int      `json:"category_id"`
}

func ParserOne(ctx *gin.Context) {
	url := ctx.Query("url")
	title := ctx.Query("menu_title")
	mid := ctx.Query("menu_id")
	id := ctx.Query("category_id")
	categoryId, _ := strconv.Atoi(id)
	mId, _ := strconv.Atoi(mid)
	RespOk(ctx, parserOne(url, title, categoryId, mId))
}
func ParserOnePron(ctx *gin.Context) {
	url := ctx.Query("url")
	title := ctx.Query("menu_title")
	mid := ctx.Query("menu_id")
	id := ctx.Query("category_id")
	categoryId, _ := strconv.Atoi(id)
	mId, _ := strconv.Atoi(mid)
	RespOk(ctx, parserOnePron(url, title, categoryId, mId))
}

func parserOne(url, title string, id, mid int) (err error) {
	if len(url) == 0 {
		return nil
	}
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)
	v := &Video{}
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})
	c.OnHTML("//input", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})
	c.OnXML("//p", func(e *colly.XMLElement) {
		//fmt.Printf("p-->>>%s\n", e.Text)
		if strings.Contains(e.Text, "：") {
			titles := strings.Split(e.Text, "：")
			if titles[0] == "片名" {
				v.Title = titles[1]
			}
			if titles[0] == "豆瓣" {
				v.Rate = titles[1]
			}
			if titles[0] == "类型" {
				v.Types = titles[1]
			}
			if titles[0] == "演员" {
				v.Actor = titles[1]
			}
			if titles[0] == "年代" {
				v.Year = titles[1]
			}
		}

	})
	c.OnHTML("font", func(e *colly.HTMLElement) {
		//baiduBtn := e.Attr("value")
		//fmt.Println("匹配到目标元素ID su:", e.Text)
		if e.Text != "全选" {
			v.URLs = append(v.URLs, e.Text)
		}
		v.Urls = strings.Join(v.URLs, " ")
	})
	c.OnHTML("img", func(e *colly.HTMLElement) {
		//baiduBtn := e.Attr("value")
		//fmt.Println("匹配到目标元素ID img:", e.Attr("src"))
		v.ThemeUrl = e.Attr("src")
	})
	c.OnHTML(".vod_content", func(e *colly.HTMLElement) {
		//fmt.Printf("vod_content-->>>%s\n", e.Text)
		v.Desc = e.Text
	})
	//c.OnResponse(func(r *colly.Response) {
	//	r.Ctx.Put("Custom-header", r.Headers.Get("Custom-Header"))
	//	fmt.Printf("copyContent-->>>%v\n", string(r.Body))
	//})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err.Error())
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	err = c.Visit(url)

	log.Printf("video -->>> %v\n", v)
	if len(v.Urls) > 0 {
		us := v.URLs[0]
		titles := strings.Split(us, "$")
		if len(titles) > 1 {
			v.Urls = titles[1]
		}
		v.MenuTitle = title
		v.MenuId = mid
		v.CategoryId = id
		v.Author = "脚本"
		b, _ := json.Marshal(&v)
		var m map[string]string
		_ = json.Unmarshal(b, &m)
		fmt.Println(m)

		resp, _ := http.Post("http://127.0.0.1:8080/api/v1/videos/insert", "application/json; charset=utf-8", bytes.NewReader(b))

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
	return
}
func parserOnePron(url, title string, id, mid int) (err error) {
	if len(url) == 0 {
		return nil
	}
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)
	v := &Video{}
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("div.detail p", func(e *colly.HTMLElement) {
		fmt.Printf("detail -->>>%s\n", e.Text)
		if strings.Contains(e.Text, "：") {
			titles := strings.Split(e.Text, "：")
			if titles[0] == "主演" {
				v.Actor = titles[1]
			}
			if titles[0] == "类型" {
				v.Types = titles[1]
			}
			if titles[0] == "备注" {
				log.Printf("备注--->>>%v\n", titles[1])
				if len(titles[1]) > 0 {
					v.Desc = titles[1]
				} else {
					v.Desc = " "
				}
			}
			if titles[0] == "上映日期" {
				v.Year = titles[1]
			}
		}
	})

	c.OnHTML("h1.limit", func(e *colly.HTMLElement) {
		fmt.Printf("title -->>>%s\n", e.Text)
		v.Title = e.Text

	})
	c.OnHTML("img#detail-img", func(e *colly.HTMLElement) {
		fmt.Printf("image -->>>%s\n", e.Attr("src"))
		v.ThemeUrl = e.Attr("src")

	})
	c.OnHTML("div.link input", func(e *colly.HTMLElement) {
		fmt.Printf("videoName-->>>%s\n", e.Attr("value"))
		if len(e.Attr("value")) > 1 {
			v.URLs = append(v.URLs, e.Attr("value"))
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err.Error())
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	log.Printf("video -->>> %v\n", v)
	if len(v.URLs) > 0 {
		us := v.URLs[0]
		titles := strings.Split(us, "$")
		if len(titles) > 1 {
			v.Urls = titles[1]
		}
		v.MenuTitle = title
		v.MenuId = mid
		v.CategoryId = id
		v.Author = "脚本"
		b, _ := json.Marshal(&v)
		var m map[string]string
		_ = json.Unmarshal(b, &m)
		fmt.Println(m)

		resp, _ := http.Post("http://127.0.0.1:8080/api/v1/videos/insert", "application/json; charset=utf-8", bytes.NewReader(b))

		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
	return
}
