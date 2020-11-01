package controllers

import (
	"fmt"
	"myweb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddArticle ...
func AddArticle(c *gin.Context) {
	//获取表单信息
	title := c.PostForm("title")
	content := c.PostForm("content")
	tags := c.PostForm("tags")
	author := c.PostForm("author")
	fmt.Println(title, content, author)

	articleModel := &models.Article{
		Title:   title,
		Content: content,
		Tags:    tags,
		Author:  author,
		Short:   content[:30],
	}

	articleID, err := articleModel.AddArticle()
	fmt.Println(articleID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "ok"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "error"})
	}
}
