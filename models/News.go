package models

type News struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"size:40"`
	Description string `json:"description" gorm:"size:200"`
	Content     string `json:"content"`
	Tag         string `json:"tag" gorm:"size:60"`
}
