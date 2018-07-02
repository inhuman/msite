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

	return r
}

func AuthRequired(c *gin.Context) {

	token := c.GetHeader("X-AUTH-TOKEN")

	_, ok := cache.GetUserTokens()[token]
	if !ok {
		c.JSON(401, gin.H{"error": "auth required"})
	}

}

func Setup2() *gin.Engine {

	var DB = make(map[string]string)
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(200, gin.H{"user": user, "value": value})
		} else {
			c.JSON(200, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.GET("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `form:"value" json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			DB[user] = json.Value
			c.JSON(200, gin.H{"status": "ok"})
		}
	})

	return r
}
