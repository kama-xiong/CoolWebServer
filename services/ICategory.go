package services

import (
	datasource "WebServer/datasources"
	"WebServer/models"
	"gorm.io/gorm"
)

type ICategoryService interface {
	GetRoot() models.Category
	GetNextLayerNode(id int) []models.Category
	GetAllChildrenNodes(id int) []models.Category
	AddRoot(node models.Category)
	AddChildNode(parentnode models.Category, childnode models.Category)
	DeleteNodes(id int)
	UpdateNode(Node models.Category)
}

type CategoryService struct {
}

func (*CategoryService) GetRoot() models.Category {
	var c models.Category
	datasource.Db.Where("layer=?").First(&c)
	return c
}
func (*CategoryService) GetNextLayerNode(id int) *[]models.Category {
	var children []models.Category
	var parent models.Category
	datasource.Db.Where("id=?", id).First(&parent)
	datasource.Db.Where("lft>? and right<? and layer=?", parent.Lft, parent.Rgt, parent.Layer+1).Find(&children)
	return &children
}
func (*CategoryService) GetAllChildrenNodes(id int) *[]models.Category {
	var children []models.Category
	var parent models.Category
	datasource.Db.Where("id=?", id).First(&parent)
	datasource.Db.Where("lft>? and right<?", parent.Lft, parent.Rgt).Find(&children)
	return &children

}
func (*CategoryService) AddRoot(node models.Category) {
	datasource.Db.Create(&node)
}
func (*CategoryService) AddChildNode(curNode models.Category, node models.Category) {
	tx := datasource.Db.Begin()
	categories := []models.Category{}
	node.Lft = curNode.Rgt
	node.Rgt = curNode.Rgt + 1
	node.Layer = curNode.Layer + 1
	/// 左码小于，右码大于当前节点的节点，右码+=2
	tx.Model(&categories).Where("lft<? and rgt>?", curNode.Lft, curNode.Rgt).Update("rgt", gorm.Expr("rgt+?", 2))

	//lft>父节点的，左右节点7均+2
	tx.Model(&categories).Where("lft>?", curNode.Rgt).Updates(map[string]interface{}{"lft": gorm.Expr("lft+?", 2), "rgt": gorm.Expr("rgt+?", 2)})

	//更新当前节点
	curNode.Rgt += 2
	tx.Save(&curNode)
	//新增节点
	tx.Create(&node)
	tx.Rollback()
	tx.Commit()

}
func (*CategoryService) DeleteNode(id int) {
	datasource.Db.Delete(&models.Category{}, id)
}
func (*CategoryService) UpdateNode(node models.Category) {
	datasource.Db.Save(&node)

}
func (*CategoryService) GetCategoriesByName(category ...string) *models.PageData {
	var boxcategory string
	var smallcategory string
	if len(category) == 1 {
		boxcategory = category[0]
	} else if len(category) == 2 {
		boxcategory = category[0]
		smallcategory = category[1]
	}
	pdata := models.PageData{}
	//获取PaperBox Category
	datasource.Db.Model(&models.Category{}).Where("name=? and layer=?", boxcategory, 1).First(&pdata.ParentCategory)
	//获取当前 Category
	datasource.Db.Model(&models.Category{}).Where("name=? and lft>? and rgt<?", smallcategory, pdata.ParentCategory.Lft, pdata.ParentCategory.Rgt).Find(&pdata.CurCategory)

	if pdata.CurCategory.Rgt-pdata.CurCategory.Lft == 1 {
		datasource.Db.Model(&models.Product{}).Where("category_id=?", pdata.CurCategory.ID).Find(&pdata.Proucts)
	} else {
		datasource.Db.Model(&models.Category{}).Where("lft>? and rgt<? and layer=?", pdata.CurCategory.Lft, pdata.CurCategory.Rgt, pdata.CurCategory.Layer+1).Find(&pdata.NavCategories)

	}
	return &pdata
}
