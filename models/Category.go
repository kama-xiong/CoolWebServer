package models

type Category struct {
	ID     int64  `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"size:40"`
	Lft    int64  `json:"lft"`
	Rgt    int64  `json:"rgt"`
	Layer  int64  `json:"level"`
	ImgUrl string `json:"imgUrl" gorm:"size:100"`
	Route  string `json:"route" gorm:"size:100"`
}

func (Category) TableName() string {
	return "Categories"
}
