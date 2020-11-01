package controllers

import (
	"fmt"
	"myweb/models"
	"myweb/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterGet ...
func RegisterGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

// RegisterPost 处理注册
func RegisterPost(c *gin.Context) {
	//获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username, password, repassword)

	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	user := &models.User{
		Username:   username,
		Status:     1,
		Createtime: time.Now().Unix(),
	}

	id, _ := user.QueryOne(username)
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名已经存在"})
		return
	}
	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	user.Password = utils.MD5(password)
	fmt.Println("md5后：", password)

	userID, err := user.AddUser()
	fmt.Println(userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册成功"})
	}
}
