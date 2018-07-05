package endpoints_test

import (
	"testing"
	"github.com/inhuman/msite/media"
	"encoding/json"
	"net/http"
	"bytes"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/inhuman/msite/cache"
	"github.com/inhuman/msite/user"
)

func TestCreatePlaylist(t *testing.T) {

	p := media.Playlist{}

	u := user.User{}
	u.ID = 1023
	u.Login = "test"
	u.Password = "test_password"

	cache.InvalidateUserTokens()
	cache.AddUserToken(&u, "c1c17c916ba11acb38b41a3d99dc678a5b3a3d78")


	jsonStr, _ := json.Marshal(p)

	req, _ := http.NewRequest("POST", "/api/user/playlist", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-AUTH-TOKEN", "c1c17c916ba11acb38b41a3d99dc678a5b3a3d78")

	w := httptest.NewRecorder()
	mockedRouter.ServeHTTP(w, req)

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
