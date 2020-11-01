package db

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	userName = "root"
	password = "mysql123"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "myweb"
)

var (
	// DB 数据库连接实例
	DB *gorm.DB
)

// InitMysql 初始化连接数据库
func InitMysql() (err error) {
	if DB == nil {
		fmt.Println("init mysql")
		path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
		DB, err = gorm.Open("mysql", path)
		if err != nil {
			return
		}
	}
	return nil
}

// func InitMysql() (err error) {
// 	if DB == nil {
// 		fmt.Println("init mysql")
// 		path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
// 		DB, err = sql.Open("mysql", path)
// 		if err != nil {
// 			return
// 		}
// 		// 设置最大连接时间
// 		DB.SetConnMaxLifetime(100 * time.Second)
// 		// 设置最大连接数
// 		DB.SetMaxIdleConns(10)
// 		// 测试连接
// 		if err = DB.Ping(); err != nil {
// 			return
// 		}
// 	}
// 	return nil
// }

// ModifyDB ...
// func ModifyDB(sql string, args ...interface{}) (int64, error) {
// 	result, err := DB.Exec(sql, args...)
// 	if err != nil {
// 		log.Println(err)
// 		return 0, err
// 	}
// 	count, err := result.RowsAffected()
// 	if err != nil {
// 		log.Println(err)
// 		return 0, err
// 	}
// 	return count, nil
// }

// QueryRowDB ...
// func QueryRowDB(sql string) *sql.Row {
// 	return DB.QueryRow(sql)
// }
