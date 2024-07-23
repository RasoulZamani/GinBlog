package models

import (
	"github.com/RasoulZamani/hiGin/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar:255"`
	Content string `gorm:"text"`
	UserID  uint
	User    models.User
}
