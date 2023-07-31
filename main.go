package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/maibokkrub/simple-backend/controller"
	model "github.com/maibokkrub/simple-backend/models"
)

func main() {
	Db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// todo: remove on production systems
	Db.AutoMigrate(&model.User{}, &model.Appointment{}, &model.Comment{})

	controller := controller.Controller{
		DB: Db,
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	appointment := r.Group("/appointment")
	{
		appointment.GET("/", controller.GetAllAppointments)
		appointment.POST("/", controller.CreateAppointment)
		appointment.PATCH("/", controller.UpdateAppointment)

		appointment.GET("/:id", controller.GetAppointmentById)
		appointment.PATCH("/archive/:id", controller.ArchiveAppointment)

		appointment.GET("/comment/:id", controller.GetAllComments)
		appointment.POST("/comment", controller.CreateComment)
	}

	user := r.Group("/user")
	{
		user.POST("", controller.CreateUser)
		user.GET("", controller.GetAllUsers)
	}

	r.Run(":8080")
}
