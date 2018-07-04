package cache

import (
	"github.com/inhuman/msite/user"
	"github.com/gin-gonic/gin"
	"errors"
)

var tokensCached = make(map[string]*user.User)

// BuildServiceTokenCache is used for build service tokens cache from db.
// Tokens cache used for authorize services
func BuildUserTokenCache() {

	usrs := user.GetAllUsers()

	for _, usr := range usrs {
		tokensCached[user.GetUserToken(&usr)] = &usr
	}
}

// GetServiceTokens is used for receive service tokens cache
func GetUserTokens() map[string]*user.User {
	return tokensCached
}

// AddServiceToken is used to add service token to cache
func AddUserToken(u *user.User, token string) {
	tokensCached[token] = u
}

// InvalidateServiceTokens is used to invalidate service tokens cache
func InvalidateUserTokens() {
	tokensCached = make(map[string]*user.User)
}

func GetCurrentUser(c *gin.Context) (*user.User, error) {

	token := c.GetHeader("X-AUTH-TOKEN")

	u, ok := GetUserTokens()[token]

	if !ok {
		return nil, errors.New("user not found")
	}

	return u, nil
}