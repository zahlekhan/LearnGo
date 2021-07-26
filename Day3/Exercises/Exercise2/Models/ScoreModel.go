package Models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	SubjectID uint
	Subject   Subject
	Marks     int `gorm:"check:(marks >=0 AND marks <=100)"";not null`
	StudentID uint
}

func (s *Score) TableName() string {
	return "score"
}
