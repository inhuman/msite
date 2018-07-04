package user

import (
	"github.com/inhuman/msite/db"
	"github.com/jinzhu/gorm"
	"crypto/sha1"
	"encoding/hex"
	"github.com/inhuman/msite/media"
)

type Register struct {
	gorm.Model
	Login           string `form:"login" json:"login" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required"`
}

type User struct {
	gorm.Model
	Login     string           `json:"login" gorm:"not null;unique"`
	Password  string           `json:"password" gorm:"not null"`
	Playlists []media.Playlist `json:"playlists"`
}

func GetUserToken(u *User) string {
	h := sha1.New()
	h.Write([]byte(u.Login + u.Password))
	return hex.EncodeToString(h.Sum(nil))
}

func GetAllUsers() []User {

	users := []User{}

	db.Stor.Db().Find(&users)

	return users
}
