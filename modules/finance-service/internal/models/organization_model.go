package models

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Id    string `json:"id"`
	GSTIn string `json:"gstIn"`
	Name  string `json:"name"`
}
