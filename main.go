package main

import (
	_ "GOYDO/src/docs"
	"GOYDO/src/modules/task/models"
	"GOYDO/src/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "host=localhost user=postgres password=dima3raza dbname=goydo port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Task{})
	if err != nil {
		panic("failed to migrate database")
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run(":3000")
	if err != nil {
		panic("failed to run server")
	}
}
