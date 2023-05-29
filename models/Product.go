package models

type Product struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"size:40"`
	PackingInfo string `json:"packingInfo" gorm:"size:60"`
	Size        string `json:"size" gorm:"size:20"`
	Cbm         string `json:"cbm" gorm:"size:20"`
	Gw          string `json:"gw" gorm:"size:20"`
	Nw          string `json:"nw" gorm:"size:20"`
	ImgUrl      string `json:"imgUrl" gorm:"size:100"`
	CategoryID  int64  `json:"categoryID"`
	Category    Category
}

func (Product) TableName() string {
	return "Products"
}
