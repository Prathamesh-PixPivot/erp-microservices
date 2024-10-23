package models

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	GSTIn string `json:"gstIn"`
	Name  string `json:"name"`
}
