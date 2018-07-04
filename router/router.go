package router

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/endpoints"
	"net/http"
	"github.com/inhuman/msite/config"
	"github.com/inhuman/msite/cache"
)

func Setup(h ...gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.Use(h...)

	// auth
	r.POST("/register", endpoints.RegisterUser)
	r.POST("/login", endpoints.LoginUser)

	secured := r.Group("/api")
	secured.Use(AuthRequired)

	// user actions
	secured.GET("/user/profile", endpoints.ProfileUser)

	secured.GET("/user/playlist", endpoints.GetPlaylists)
	secured.POST("/user/playlist", endpoints.CreatePlaylist)
	secured.POST("/user/playlist/media", endpoints.AddMediaToPlaylist)
	secured.DELETE("/user/playlist/media", endpoints.RemoveMediaFromPlaylist)
	secured.DELETE("/user/playlist", endpoints.DeletePlaylist)

	// media
	secured.POST("/media/upload", endpoints.UploadFile)
	secured.StaticFS("/media/", http.Dir(config.AppConf.UploadPath))

	return r
}

func AuthRequired(c *gin.Context) {

	_, err := cache.GetCurrentUser(c)

	if err != nil {
		c.JSON(401, gin.H{"error": "auth required"})
		c.Abort()
	}

}
