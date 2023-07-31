package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maibokkrub/simple-backend/dto"
	model "github.com/maibokkrub/simple-backend/models"
)

func (controller *Controller) CreateAppointment(c *gin.Context) {
	var dto dto.CreateAppointmentDTO

	if err := c.BindJSON(&dto); err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	appointment, err := dto.ToModel()
	if err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId := c.GetInt("userID")
	appointment.CreatedBy = userId

	appointment.Create(controller.DB)
}

func (controller *Controller) GetAllAppointments(c *gin.Context) {
	pageQuery := c.DefaultQuery("page", "0")

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := model.GetAllAppointment(controller.DB, page)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *Controller) GetAppointmentById(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil || id < 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := model.GetOneAppointmentWithComments(controller.DB, int(id))
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if result == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (controller *Controller) UpdateAppointment(c *gin.Context) {
	var dto dto.UpdateAppointmentDTO

	if err := c.BindJSON(&dto); err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	oldData, err := model.GetOneAppointment(controller.DB, dto.ID)
	if err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	appointment, err := dto.ToModel(oldData)
	if err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	appointment.Update(controller.DB)
}

func (controller *Controller) ArchiveAppointment(c *gin.Context) {
	idQuery := c.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil || id < 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	oldData, err := model.GetOneAppointment(controller.DB, int(id))
	if err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := oldData.SoftDelete(controller.DB)
	c.JSON(200, res.RowsAffected)
}
