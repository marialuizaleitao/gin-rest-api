package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name string `json:"name" validate:"required,min=2,max=40"`
	Role string `json:"role" validate:"required,min=2,max=100"`
}

func ValidateMemberData(member *Member) error {
	if err := validator.Validate(member); err != nil {
		return err
	}
	return nil
}
