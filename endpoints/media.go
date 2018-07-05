package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/config"
	"github.com/inhuman/msite/media"
	"github.com/inhuman/msite/db"
	"github.com/inhuman/msite/cache"
	"strconv"
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

	playlist := &media.Playlist{}

	err := c.Bind(playlist)
	if err != nil {
		c.JSON(402, err)
		return
	}

	u, _ := cache.GetCurrentUser(c)

	playlist.UserID = u.ID

	db.Stor.Db().Save(playlist)
}

func GetPlaylists(c *gin.Context) {

	u, _ := cache.GetCurrentUser(c)

	p := []media.Playlist{}

	db.Stor.Db().Model(&u).Related(&p)

	c.JSON(200, p)

}

func DeletePlaylist(c *gin.Context) {

	u, _ := cache.GetCurrentUser(c)

	playlistId, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(402, err)
		return
	}

	p := media.Playlist{}
	p.ID = uint(playlistId)

	r := db.Stor.Db().Model(&u).Related(&p).Unscoped().Delete(&p).RowsAffected

	if r == 0 {
		c.JSON(404, gin.H{"error": "playlist not found"})
	}

}

func AddMediaToPlaylist(c *gin.Context) {

	playlist := &media.Playlist{}

	err := c.Bind(playlist)
	if err != nil {
		c.JSON(402, err)
		return
	}

	// find playlist
	// ensure that playlist belongs to user
	// add given media to playlist
	// save playlist


}

func RemoveMediaFromPlaylist(c *gin.Context) {
	// find playlist
	// ensure that playlist belongs to user
	// remove given media to playlist
	// save playlist
}