package models

import (
	"gopkg.in/validator.v2"
	"time"
)

type Member struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `json:"name" validate:"required,min=2,max=40"`
	Role      string     `json:"role" validate:"required,min=2,max=100"`
}

func ValidateMemberData(member *Member) error {
	if err := validator.Validate(member); err != nil {
		return err
	}
	return nil
}
