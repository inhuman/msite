package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/config"
)

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, err)
	}

	err = c.SaveUploadedFile(file, config.AppConf.UploadPath + "/" + file.Filename)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{"media": gin.H{"url": "/api/media/" + file.Filename}})
}

func CreatePlaylist(c *gin.Context) {

	//db.Stor.Db().Model(user.User{}).Related(media.Playlist{})

}

func GetPlaylists(c *gin.Context) {

}

func DeletePlaylist(c *gin.Context) {

}

func UpdatePlaylist(c *gin.Context) {

}