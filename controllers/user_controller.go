package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc UserController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": nil})
}
