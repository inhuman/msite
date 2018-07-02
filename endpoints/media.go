package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/config"
)

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(500, err)
	}

	err = c.SaveUploadedFile(file, config.AppConf.UploadPath + "/" + file.Filename)
	if err != nil {
		c.JSON(500, err)
	}

	//TODO: generate url
}
