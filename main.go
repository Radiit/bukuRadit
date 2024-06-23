package main

import (
	"bukuRadit/db"
	"bukuRadit/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := gin.Default()

	routes.BookRoute(router)

	router.Run()
}
