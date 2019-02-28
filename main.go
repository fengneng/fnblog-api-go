package main

import (
	"faBlog-api-go/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cgo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// db.DropTableIfExists(&entity.Topic{})
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&entity.Topic{})
	// // 创建
	// db.Create(&entity.Topic{Content: "如果成为一个厉害的工程师", Author: "李子木"})

	// // 读取
	// var topic entity.Topic
	// db.First(&topic, 1) // 查询id为1的product

	router := gin.Default()

	RegisterWebPage(router)
	RegisterMobileApi(router)
	RegisterWebApi(db, router)

	router.Run(":8080")
}
