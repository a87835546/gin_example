package middleware

import (
	"gin_example/controllers"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"log"
)

func Auth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		// 获取方法
		act := c.Request.Method
		sub, _ := c.Get("rule") // "root"

		// 判断策略是否已经存在了
		if ok, _ := e.Enforce(sub, obj, act); ok {
			log.Println("Check successfully")
			c.Next()
		} else {
			log.Println("sorry , Check failed")
			controllers.RespAuthError(c)
			c.Abort()
		}
	}
}
