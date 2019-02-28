package entity

import (
	"github.com/jinzhu/gorm"
)

type Topic struct {
	gorm.Model
	Content string
	Author  string
}
