package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlserver/application"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

type urlReq struct {
	Url  string `json:"url" binding:"required"`
	Time int64  `json:"time" binding:"required"`
}

func main() {
	r := gin.Default()
	r.Use(Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/url/toS", func(context *gin.Context) {
		var pram urlReq
		context.ShouldBindJSON(&pram)
		if pram.Url == "" {
			context.JSON(500, gin.H{
				"url":   nil,
				"error": "参数url为空",
			})
			return
		}
		shortUrl := application.LongToShort(pram.Url, pram.Time)
		context.JSON(200, gin.H{
			"url": shortUrl,
		})
	})
	r.GET("/go/:id", func(context *gin.Context) {
		id := context.Param("id")
		lUrl := application.ShortToLong(id)
		context.Redirect(http.StatusMovedPermanently, lUrl)
	})
	r.Run(":8811")
}
