package main

import (
	"GOYDO/src/cfg"
	_ "GOYDO/src/docs"
	"GOYDO/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r, cfg.DB)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run(":3000")
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
