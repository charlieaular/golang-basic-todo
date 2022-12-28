package main

import (
	"golang-todo/src/database"
	"golang-todo/src/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	mysqlIns := database.MysqlConn

	db := mysqlIns.Init()

	r := gin.Default()

	r.Use(cors.Default())

	routes.RegisteTodoRoutes(r, db)

	r.Run(":9000")
}
