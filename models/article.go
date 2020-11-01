package models

import "myweb/db"

// Article 文章模型
type Article struct {
	ID         int    `gorm:"id"`
	Title      string `gorm:"title"`
	Tags       string `gorm:"tags"`
	Short      string `gorm:"short"`
	Content    string `gorm:"content"`
	Author     string `gorm:"author"`
	CreateTime int64  `gorm:"create_time"`
	UpdateTime int64  `gorm:"updata_time"`
	Status     int    `gorm:"status"` //Status=0为正常，1为删除，2为冻结
}

// TableName ...
func (Article) TableName() string {
	return "articles"
}

// AddArticle ...
func (a *Article) AddArticle() (int, error) {
	articleDB := db.DB
	articleDB.Create(a)
	return 1, nil
}

// QueryOneWithID ...
func (a *Article) QueryOneWithID(articleID int) *Article {
	article := &Article{}
	articleDB := db.DB
	articleDB.Where("id = ?", articleID).Find(article)
	return article
}
