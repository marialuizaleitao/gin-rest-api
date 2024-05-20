package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name string `json:"name"`
	Role string `json:"role"`
}
