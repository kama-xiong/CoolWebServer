package models

type PageData struct {
	ParentCategory Category
	CurCategory    Category
	NavCategories  []Category
	Proucts        []Product
	PageNumber     int
	PageSize       int
	TotalCount     int
}
