package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterMobileApi(router *gin.Engine) {

	router.GET("/login", login)

	v1 := router.Group("api/v1")

	v1.GET("/hot_rcmd", func(c *gin.Context) {
		jsonFile(c, "hot_rcmd.json")
	})
	v1.GET("/channel", func(c *gin.Context) {
		jsonFile(c, "channel.json")
	})
	v1.GET("/channel_detail", func(c *gin.Context) {
		jsonFile(c, "channel_detail.json")
	})
	v1.GET("/topic_detail", func(c *gin.Context) {
		jsonFile(c, "topic_detail.json")
	})
	v1.GET("/article_detail", func(c *gin.Context) {
		jsonFile(c, "article_detail.json")
	})
	v1.GET("/attention", func(c *gin.Context) {
		jsonFile(c, "attention.json")
	})
	v1.GET("/follows", func(c *gin.Context) {
		jsonFile(c, "follows.json")
	})
	v1.GET("/photo_flow", func(c *gin.Context) {
		jsonFile(c, "photo_flow.json")
	})
	v1.GET("/message", func(c *gin.Context) {
		jsonFile(c, "message.json")
	})
	v1.GET("/user_profile", func(c *gin.Context) {
		jsonFile(c, "user_profile1.json")
	})
}

func jsonFile(c *gin.Context, file string) {
	info, err := ioutil.ReadFile("resources/dataFile/" + file)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusForbidden, "not found")
	} else {

		fmt.Println(info)
		result := string(info)

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(200, result)
	}
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}

func go404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", `Sorry,I lost myself!`)
}
