package controllers

import (
	"log"
	"start/config"
	"start/datasources/myminio"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
}

func (item *CustomerController) CheckRequest(c *gin.Context) {

	c.JSON(200, gin.H{"message": "Hello world"})
	c.Abort()
	return
}

func (item *CustomerController) Upload(c *gin.Context) {
	file, _, errAvatar := c.Request.FormFile("avatar")

	if errAvatar != nil {
		log.Println(errAvatar)
		c.JSON(400, gin.H{"message": errAvatar.Error()})
		c.Abort()
		return
	}

	config := config.GetConfig()

	fileName, errUpload := myminio.UploadFile(
		&file,
		config.GetString("minio.bucket"),
		"a.png")
	if errUpload != nil {
		log.Println("ERROR_UPLOAD:", errUpload)
		c.JSON(400, gin.H{"message": errUpload.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": fileName})
	c.Abort()
	return

}
