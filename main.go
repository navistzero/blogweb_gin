package main

import (
	"myweb/db"
	"myweb/routes"
)

func main() {
	db.InitMysql()
	router := routes.InitRouter()
	router.Static("/static", "./static")
	router.Run(":8090")
}
