package media

import (
	"github.com/jinzhu/gorm"
)

type Playlist struct {
	gorm.Model
	Media []Media `json:"media" gorm:"many2many:playlist_medias;association_autoupdate:false;"`
	UserID uint `json:"user_id"`
}
