package Models

import (
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name string `not null`
}

func (s *Subject) TableName() string {
	return "subject"
}
