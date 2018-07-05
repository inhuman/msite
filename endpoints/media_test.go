package endpoints_test

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/inhuman/msite/router"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/inhuman/msite/db"
	"github.com/inhuman/msite/media"
	"encoding/json"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestCreatePlaylist(t *testing.T) {

	mocket.Catcher.Register()
	dbm, err := gorm.Open(mocket.DRIVER_NAME, "any_string")

	if err != nil {
		t.Fatal(err)
	}

	db.Stor.SetDb(dbm)

	gin.SetMode(gin.TestMode)
	r := router.Setup()

	p := media.Playlist{}

	jsonStr, _ := json.Marshal(p)

	req, _ := http.NewRequest("POST", "/api/user/playlist", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeletePlaylist(t *testing.T) {

}

func TestGetPlaylists(t *testing.T) {

}

func TestAddMediaToPlaylist(t *testing.T) {

}

func TestRemoveMediaFromPlaylist(t *testing.T) {

}

func TestUploadFile(t *testing.T) {

}
