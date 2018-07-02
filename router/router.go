package router

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/cache"
	"github.com/inhuman/msite/endpoints"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	r.POST("/register", endpoints.RegisterUser)
	r.POST("/login", endpoints.LoginUser)

	secured := r.Group("/api")
	secured.Use(AuthRequired)

	secured.GET("/user/profile", endpoints.ProfileUser)

	secured.POST("/media/upload", endpoints.UploadFile)


	//TODO: add url to serve folder

	return r
}

func AuthRequired(c *gin.Context) {

	token := c.GetHeader("X-AUTH-TOKEN")

	_, ok := cache.GetUserTokens()[token]
	if !ok {
		c.JSON(401, gin.H{"error": "auth required"})
	}

}
