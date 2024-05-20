package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
