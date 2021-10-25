package models

import (
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}

func NewSection(title string, user_id string) *Section {
	return &Section{
		Title:  title,
		UserId: user_id,
	}
}

func (s *Section) Create() *Section {
	section := &Section{Title: s.Title, UserId: s.UserId}
	DbConnection.Create(&section)
	return section
}
