package models

import (
	"myweb/db"
	"myweb/entity"

	"github.com/jinzhu/gorm"
)

// ArticleModel 文章模型
type ArticleModel struct {
	// Dao *models.Dao
	db *gorm.DB
}

// NewArticleModel 文章模型new方法
func NewArticleModel() *ArticleModel {
	return &ArticleModel{
		db: db.DB,
	}
}

// Add ...
func (a *ArticleModel) Add(data *entity.ArticleEntity) (int, error) {
	a.db.Create(data)
	// dataRes := a.db.Create(a)
	return 1, nil
}

// // QueryOneWithID ...
// func (a *Article) QueryOneWithID(articleID int) *Article {
// 	article := &Article{}
// 	articleDB := db.DB
// 	articleDB.Where("id = ?", articleID).Find(article)
// 	return article
// }
