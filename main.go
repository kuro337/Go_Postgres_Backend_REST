package main

import (
	"main/database"
	"main/routes"
"github.com/gin-gonic/gin"
"github.com/gin-contrib/cors"
"time"
)

func main() {


database.Connect()

	router := gin.Default()
	routes.Setup(router)

	// CORS so cookies works

	router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"},
    AllowMethods:     []string{"PUT", "PATCH"},
    AllowHeaders:     []string{"Origin"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
  }))

	router.Run("localhost:8080")
}

