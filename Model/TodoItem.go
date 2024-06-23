package Model

import "github.com/jinzhu/gorm"

type TodoItem struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null;"`
	Description string `gorm:"type:varchar(255);not null;"`
	Done        bool   `gorm:"type:bool;not null;"`
}
