package main

import (
	"github.com/inhuman/msite/router"
	"github.com/inhuman/msite/db"
	"github.com/inhuman/msite/user"
	"os"
	"github.com/inhuman/msite/config"
	"log"
	"github.com/inhuman/msite/cache"
	"github.com/inhuman/msite/media"
	"github.com/gin-gonic/gin"
)

func main() {
	err := runApp()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func runApp() error {

	if len(os.Args) > 1 {
		err := config.AppConf.Load(os.Args...)
		if err != nil {
			return err
		}
	} else {
		err := config.AppConf.Load()
		if err != nil {
			return err
		}
	}

	db.Init()
	cache.BuildUserTokenCache()

	db.Stor.Migrate(user.User{})
	db.Stor.Migrate(media.Media{})
	db.Stor.Migrate(media.Playlist{})




	r := router.Setup(gin.Logger(), gin.Recovery())

	r.Run(":8080")
	return nil
}
