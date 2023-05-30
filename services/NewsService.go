package services

import (
	datasource "WebServer/datasources"
	"WebServer/models"
)

type NewsService struct {
}

func (*NewsService) GetArticle(id int) *models.News {
	article := models.News{}
	datasource.Db.First(&article, id)
	return &article
}
func (*NewsService) GetArticles(limit int) *[]models.News {
	articles := []models.News{}
	datasource.Db.Limit(limit).Find(&articles)
	return &articles
}
func (*NewsService) GetPageArticles(pagination models.Pagination) *[]models.News {
	news := []models.News{}
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	switch {
	case pagination.Pagesize > 100:
		pagination.Pagesize = 100
	case pagination.Pagesize <= 0:
		pagination.Pagesize = 10
	}
	offset := (pagination.Page - 1) * pagination.Pagesize
	datasource.Db.Offset(offset).Limit(pagination.Pagesize).Find(&news)
	return &news
}
