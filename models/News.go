package models

import "time"

type News struct {
	Id          int64     `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"size:40"`
	Author      string    `json:"author" gorm:"size:20"`
	Description string    `json:"description" gorm:"size:200"`
	Content     string    `json:"content"`
	Tag         string    `json:"tag" gorm:"size:60"`
	CreatedAt   time.Time `json:"createdat"`
}
