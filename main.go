package main

import (
	"myweb/db"
	"myweb/routes"
)

func main() {
	err := db.InitMysql()
	if err != nil {
		return
	}
	router := routes.InitRouter()
	router.Static("/static", "./static")
	router.Run(":8090")
}
