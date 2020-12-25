package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hillfolk/go-ddd-start/models"
	"net/http"
)

type UserController struct{}

var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {
	if c.Param("id") != "" {
		user:= userModel.GetByID(c.Param("id"))
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
