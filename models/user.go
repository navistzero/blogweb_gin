package models

import (
	"myweb/db"
)

// User ...
type User struct {
	ID         int    `gorm:"id"`
	Username   string `gorm:"username"`
	Password   string `gorm:"password"`
	Status     int    `gorm:"status"`
	Createtime int64  `gorm:"createtime"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}

// AddUser ...
func (u *User) AddUser() (int64, error) {
	userDB := db.DB
	userDB.Create(u)
	return 1, nil
}

// QueryOne ...
func (u *User) QueryOne(userName string) (int, error) {
	user := &User{}
	userDB := db.DB
	userDB.Where("username = ?", userName).Find(user)
	return user.ID, nil
}

//InsertUser 插入
// func InsertUser(user User) (int64, error) {
// 	return db.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
// 		user.Username, user.Password, user.Status, user.Createtime)
// }

// //QueryUserWightCon 按条件查询
// func QueryUserWightCon(con string) int {
// 	sql := fmt.Sprintf("select id from users %s", con)
// 	fmt.Println(sql)
// 	row := db.QueryRowDB(sql)
// 	id := 0
// 	row.Scan(&id)
// 	return id
// }

// //QueryUserWithUsername 根据用户名查询id
// func QueryUserWithUsername(username string) int {
// 	sql := fmt.Sprintf("where username='%s'", username)
// 	return QueryUserWightCon(sql)
// }

/*
//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string)int{
    sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
    return QueryUserWightCon(sql)
}

*/
