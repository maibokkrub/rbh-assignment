package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/maibokkrub/simple-backend/dto"
	model "github.com/maibokkrub/simple-backend/models"
)

func (controller *Controller) CreateComment(c *gin.Context) {
	var dto dto.CreateCommentDTO
	if err := c.BindJSON(&dto); err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	comment, err := dto.ToModel()
	if err != nil {
		// todo: cleanup message
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	comment.Create(controller.DB)
}

func (controller *Controller) GetAllComments(c *gin.Context) {
	idQuery := c.Param("id")

	id, err := strconv.Atoi(idQuery)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := model.GetAllComment(controller.DB, id)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
