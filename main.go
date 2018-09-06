package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// 跨域中间件 https://blog.csdn.net/xuedapeng/article/details/79076704
func Kuayu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type,XFILENAME,XFILECATEGORY,XFILESIZE, X-Token")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.JSON(204, gin.H{"code": 20000})
		}
		c.Next()
	}
}

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	r.Use(Kuayu())

	// 登录api  vue-admin-template.
	//r.OPTIONS("/api/user/login", func(c *gin.Context) {
	//	c.JSON(204, gin.H{"code": 20000})
	//})
	r.POST("/api/user/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 20000, "data": gin.H{"token": "admin"}})
		//var msg struct {
		//	Name string `json:"username" binding:"required"`
		//	Pwd  string `json:"password" binding:"required"`
		//}
		//if err := c.ShouldBindJSON(&msg); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"code": "20001", "msg": "query params error"})
		//	return
		//} else if msg.Name == "admin" && msg.Pwd == "admin" {
		//	// {"code":20000,"data":{"token":"admin"}}
		//	c.JSON(http.StatusOK, gin.H{"code": "20000", "data": gin.H{"token": "admin"}})
		//} else {
		//	c.JSON(http.StatusBadRequest, gin.H{"code": "20001", "msg": "username or password wrong"})
		//}

	})

	// 退出登录api  vue-admin-template.
	r.POST("/api/user/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 20000, "data": gin.H{"token": "admin"}})
	})

	// 获取用户信息api  vue-admin-template.
	//r.OPTIONS("/api/user/info", func(c *gin.Context) {
	//	c.JSON(204, gin.H{"code": 20000})
	//})
	r.GET("/api/user/info", func(c *gin.Context) {
		token := c.Query("token")
		if token == "admin" {
			c.JSON(http.StatusOK,
				gin.H{"code": 20000,
					"data": gin.H{"roles": []string{"admin"},
						"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"}})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": 20001, "msg": "wrong token"})
		}

	})

	// 获取table信息  vue-admin-template.
	r.GET("/api/table/list", func(c *gin.Context) {
		bytes := []byte(`[
      {"id":"630000199908276858","title":"Mvevgnajv narjnnrs nhclvd syznia klyf zmkdlc rltksrlqx mkzx wtyw jsjghhc gzzfs.","status":"deleted","author":"name","display_time":"2012-06-12 14:15:24","pageviews":2506},
      {"id":"340000198205066762","title":"Sslcw elzufuf tndiowtae ocuxp sbnbgu fnlwxwkgpv wnfzv nehxifgbpi iouqnne xvvrdf.","status":"published","author":"name","display_time":"1993-07-12 18:38:50","pageviews":3599},
      {"id":"370000201809025865","title":"Ygge knjfooxca xuiv njmlkjocuv lsvakhlvrw htet wsg glfp ofwm pls rekcctes ytjcj lpmdw.","status":"published","author":"name","display_time":"1970-09-06 05:42:28","pageviews":2869},
      {"id":"640000197710018741","title":"Kfr ifvhzjjdti vtjtdr onmkkbzm seihrhfb hprbaxm bjpgdw exjuxxk ymokdlfns cspiugnh podeoumv gewkokdxxv jsnmal klmkyxvwp gquu yuark mtipwuwh pmpxpjsket.","status":"deleted","author":"name","display_time":"1985-11-16 08:36:05","pageviews":4040},
      {"id":"430000201404214832","title":"Nfttinslvx bnpdcbgnv ilfggd vhuyydlon orzgl gndiyztw qcopjix rsoi mpvivml lmxsrjk ebluymh.","status":"deleted","author":"name","display_time":"1997-10-05 18:18:19","pageviews":4619},
      {"id":"360000197707165946","title":"Ozitv dpa urvbmlt fhj pcgclrcd qbbylpdiq mxyj ovgs enrbxquidm jxvyf tsbj mlksksxkc.","status":"draft","author":"name","display_time":"1990-05-02 23:31:44","pageviews":4189}
    ]`)

		type st struct {
			Id          string `json:"id"`
			Title       string `json:"title"`
			Status      string `json:"status"`
			Author      string `json:"author"`
			DisplayTime string `json:"display_time"`
			Pageviews   int    `json:"pageviews"`
		}

		var stArr []st
		err := json.Unmarshal(bytes, &stArr)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 20000, "data": gin.H{"items": stArr}})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": 20001, "msg": "json format error"})
		}

	})

	r.Run(":9999")
}
