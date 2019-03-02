package controllers

import (
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
}

func (item *CustomerController) CheckRequest(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Hello world"})
	c.Abort()
	return
}
