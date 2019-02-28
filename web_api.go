package main

import (
	"faBlog-api-go/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegisterWebApi(db *gorm.DB, router *gin.Engine) {

	v1 := router.Group("web_api/v1")

	v1.GET("/topics", func(c *gin.Context) {
		var topics []entity.Topic
		db.Find(&topics)
		c.JSON(http.StatusOK, topics)
	})
}
