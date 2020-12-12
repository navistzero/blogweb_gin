package controllers

import (
	"fmt"
	"myweb/entity"
	"myweb/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddArticle 新增文章
func AddArticle(c *gin.Context) {
	//获取表单信息
	title := c.PostForm("title")
	content := c.PostForm("content")
	tags := c.PostForm("tags")
	author := c.PostForm("author")
	fmt.Println(title, content, author)

	articleData := &entity.ArticleEntity{
		Title:   title,
		Content: content,
		Tags:    tags,
		Author:  author,
		Short:   content[:30],
	}

	articleServer := service.NewArticleServer()
	articleID, err := articleServer.AddArticle(articleData)
	println(articleID, err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok", "data": articleID})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "error"})
	}
}
