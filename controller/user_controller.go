package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/maibokkrub/simple-backend/models"
)

func (controller *Controller) CreateUser(c *gin.Context) {
	user := model.User{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// todo: clean up error message leak
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		// todo: clean up error message leak
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user.Create(controller.DB)
}

func (controller *Controller) GetAllUsers(c *gin.Context) {
	result, err := model.GetAllUsers(controller.DB, 0)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}
