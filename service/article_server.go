package service

import (
	"myweb/entity"
	"myweb/models"
)

// ArticleServer 文章服务对象
type ArticleServer struct {
	Cache string
	Model *models.ArticleModel
}

// NewArticleServer 创建文章服务工厂方法
func NewArticleServer() *ArticleServer {
	return &ArticleServer{
		Cache: "yes",
		Model: models.NewArticleModel(),
	}
}

// AddArticle 新增文章数据
func (a *ArticleServer) AddArticle(data *entity.ArticleEntity) (int, error) {
	return a.Model.Add(data)
}
