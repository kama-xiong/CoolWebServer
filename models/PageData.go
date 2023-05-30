package models

type PageData struct {
	ParentCategory Category
	CurCategory    Category
	NavCategories  []Category
	Products       []Product
	PageNumber     int
	PageSize       int
	TotalCount     int
}
