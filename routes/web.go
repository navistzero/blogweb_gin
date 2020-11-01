package routes

import (
	"myweb/app/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/ping", controllers.Ping)
	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	article := router.Group("/article")
	{
		//写文章
		article.POST("/add", controllers.AddArticle)
		//显示文章内容
		// article.GET("/show/:id", controllers.ShowArticleGet)
		// //更新文章
		// article.POST("/update",controllers.UpdateArticlePost)
		// //删除文章
		// article.GET("/delete",controllers.DeleteArticleGet)

	}
	return router
}
