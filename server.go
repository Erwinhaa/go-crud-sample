package main

import (
	"os"

	"myapp/config"
	"myapp/routes"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	PORT := os.Getenv("APP_PORT")

	if PORT == "" {
		PORT = defaultPort
	}

	router := gin.New()

	config.ConnectDB()
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	routes.WebRouter(router)
	router.Run(":" + PORT)
}
