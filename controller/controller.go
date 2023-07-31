package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maibokkrub/simple-backend/dto"
	"github.com/maibokkrub/simple-backend/middleware"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (controller *Controller) InitRoutes(parent *gin.RouterGroup) {
	appointment := parent.Group("/appointment")
	{
		appointment.GET("/", controller.GetAllAppointments)
		appointment.POST("/", controller.CreateAppointment)
		appointment.PATCH("/", controller.UpdateAppointment)

		appointment.GET("/:id", controller.GetAppointmentById)
		appointment.PATCH("/archive/:id", controller.ArchiveAppointment)

		appointment.GET("/comment/:id", controller.GetAllComments)
		appointment.POST("/comment", controller.CreateComment)
	}

	user := parent.Group("/user")
	{
		user.GET("/", controller.GetAllUsers)
		user.POST("/", controller.CreateUser)
	}
}

func (controller *Controller) FakeLogin(c *gin.Context) {
	var dto dto.FakeLoginDTO
	if err := c.BindJSON(&dto); err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := middleware.NewToken(dto.ID)
	if err != nil {
		// todo: cleanup message
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, token)
}

func (controller *Controller) RenewToken(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newToken, err := middleware.NewToken(userID.(int))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, newToken)
}
