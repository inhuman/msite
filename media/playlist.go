package media

import (
	"github.com/jinzhu/gorm"
)

type Playlist struct {
	gorm.Model
	Media []Media
	UserID uint
}
