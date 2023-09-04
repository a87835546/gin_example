package controllers

import (
	"fmt"
	"gin_example/doreamon"
	"gin_example/logic"
	"gin_example/models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
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
	//ctx.JSON(200, &Result{code: 200, message: "success", data: data})
	ctx.JSON(200, gin.H{"code": 200, "message": "success", "data": data})
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
func generateToken(c *gin.Context, user *models.User) {
	j := &doreamon.JWT{
		SigningKey: []byte("newtrekWang"),
	}

	claims := doreamon.CustomClaims{
		ID:   user.Id,
		Name: user.Username,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
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
	RespOk(c, gin.H{
		"code":    200,
		"message": "登录成功！",
		"data":    user,
	})
	return
}

func SystemConfig(ctx *gin.Context) {
	RespOk(ctx, nil)
}
