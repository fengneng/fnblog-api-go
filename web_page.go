package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebPage(router *gin.Engine) {
	router.LoadHTMLGlob("resources/views/**/*")
	router.NoRoute(go404)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	router.GET("index.html", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
}
