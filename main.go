package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/maibokkrub/simple-backend/controller"
	"github.com/maibokkrub/simple-backend/middleware"
	model "github.com/maibokkrub/simple-backend/models"
)

func main() {

	// use postgres for prod, maybe you make an env flag later
	//
	// config, err := GetConfigFromEnv()
	// if err != nil {
	// 	panic(err)
	// }
	// Db, err := gorm.Open(postgres.Open(config.DSN))
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// use sql lite for dev
	Db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// todo: remove on production systems
	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Appointment{})
	Db.AutoMigrate(&model.Comment{})

	controller := controller.Controller{
		DB: Db,
	}

	// todo: init CORS middleware
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.POST("/login", controller.FakeLogin)

	api := r.Group("/api", middleware.AuthMiddleware())
	api.GET("/renew", controller.RenewToken)

	v1 := api.Group("/v1")
	controller.InitRoutes(v1)

	r.Run(":8080")
}
