package models

import (
	"time"
)

type Section struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	UserId    string    `json:"user_id"`
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

func GetSections(user_id string) []Section {
	var sections []Section
	DbConnection.Where("user_id = ?", user_id).Find(&sections)
	return sections
}
