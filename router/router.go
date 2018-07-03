package router

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/cache"
	"github.com/inhuman/msite/endpoints"
	"net/http"
	"github.com/inhuman/msite/config"
)

func Setup(h ...gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.Use(h...)

	r.POST("/register", endpoints.RegisterUser)
	r.POST("/login", endpoints.LoginUser)

	secured := r.Group("/api")
	secured.Use(AuthRequired)

	secured.GET("/user/profile", endpoints.ProfileUser)

	secured.POST("/user/playlist")


	secured.POST("/media/upload", endpoints.UploadFile)
	secured.StaticFS("/media/", http.Dir(config.AppConf.UploadPath))

	return r
}

func AuthRequired(c *gin.Context) {

	token := c.GetHeader("X-AUTH-TOKEN")

	_, ok := cache.GetUserTokens()[token]
	if !ok {
		c.JSON(401, gin.H{"error": "auth required"})
		c.Abort()
	}

}
