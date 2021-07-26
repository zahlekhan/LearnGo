package Models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string
	LastName string
	DOB      string
	Address  string
	Scores   []Score
}

func (s *Student) TableName() string {
	return "student"
}
