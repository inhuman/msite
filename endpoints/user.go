package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/cache"
	"fmt"
	"github.com/inhuman/msite/user"
	"github.com/inhuman/msite/db"
)

func RegisterUser(c *gin.Context) {

	ur := &user.Register{}

	err := c.Bind(ur)

	if err != nil {
		c.JSON(402, err)
		return
	}

	u := &user.User{
		Login:    ur.Login,
		Password: ur.Password,
	}

	dbErr := db.Stor.Db().Save(u).Error
	if dbErr != nil {
		c.JSON(500, dbErr)
		return
	}
	cache.AddUserToken(u, user.GetUserToken(u))

	u.Password = "********"
	c.JSON(200, u)
}

func LoginUser(c *gin.Context)  {

	u := &user.User{}

	err := c.Bind(u)

	if err != nil {
		c.JSON(402, err)
		return
	}

	fmt.Printf("%+v\n", u)

	res := db.Stor.Db().Where(&user.User{Login: u.Login, Password: u.Password}).First(u)

	if res.RowsAffected == 1 {
		fmt.Printf("user found %+v\n", u)
	}

	switch  {
	case res.RowsAffected == 0:
		c.JSON(404, gin.H{"error": "user not found"})
	case res.RowsAffected == 1:

		c.JSON(200, gin.H{"token": user.GetUserToken(u)})
	case res.RowsAffected > 1:
		c.JSON(500, gin.H{"error": "user collision"})
	}

}

func ProfileUser(c *gin.Context) {

	token := c.GetHeader("X-AUTH-TOKEN")

	usr := cache.GetUserTokens()[token]

	c.JSON(200, usr)

}