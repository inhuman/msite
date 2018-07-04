package media

import "github.com/jinzhu/gorm"

type Media struct {
	gorm.Model
	Uri string `json:"uri" gorm:"unique"`
}

