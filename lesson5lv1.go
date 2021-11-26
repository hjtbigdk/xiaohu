package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
//c.Next() 调用后续的处理函数 c.Abort()阻止调用后续函数
func main() {
	r := gin.Default()
	auth := func(c *gin.Context) {
		value, err := c.Cookie("gin_cookie")
		if err != nil {
			c.JSON(403, gin.H{
				"message": "认证失败,没有cookie",
			})
			c.Abort()
		} else {
			c.Set("cookie", value)
			c.Next()
			v, _ := c.Get("next")
			fmt.Println(v)
		}
	}

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "123" && password == "321" {//”==“给key”username“的value赋值
			c.SetCookie("gin_cookie", username, 3600, "/", "", false, true)
			c.JSON(200, gin.H{
				"msg": "login successfully",
			})
		} else {
			c.JSON(403, gin.H{
				"message": "认证失败,账号密码错误",
			})
		}
	})

	r.GET("/hello", auth, func(c *gin.Context) {
		cookie, _ := c.Get("cookie")
		str := cookie.(string)
		c.String(200, "hello world"+str)
		c.Set("next", "test next")
	})

	r.Run()
}