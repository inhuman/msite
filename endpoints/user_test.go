package endpoints_test

import (
	"testing"
	"net/http/httptest"
	"github.com/inhuman/msite/router"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"bytes"
	"github.com/inhuman/msite/user"
	"encoding/json"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/inhuman/msite/cache"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/inhuman/msite/db"
)

var mockedRouter *gin.Engine

func init(){
	mocket.Catcher.Register()
	dbm, err := gorm.Open(mocket.DRIVER_NAME, "any_string")

	if err != nil {
		log.Fatal(err)
	}

	db.Stor.SetDb(dbm)

	gin.SetMode(gin.TestMode)
	mockedRouter = router.Setup()
}

func TestRegisterUser(t *testing.T) {

	ur := &user.Register{
		Login: "test",
		Password: "test_password",
		ConfirmPassword: "test_password",
	}

	mocket.Catcher.NewMock().WithQuery(`INSERT INTO "users" (.+)`)

	jsonStr, _ := json.Marshal(ur)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mockedRouter.ServeHTTP(w, req)

	u := &user.User{}

	json.Unmarshal(w.Body.Bytes(), u)

	assert.Equal(t, u.Login, "test")
	assert.Equal(t, u.Password, "********")
}


func TestLoginUser(t *testing.T) {

	u := &user.User{
		Login: "test",
		Password: "test_password",
	}
	jsonStr, _ := json.Marshal(u)

	commonReply := []map[string]interface{}{{"login": "test", "password": "test_password"}}

	mocket.Catcher.NewMock().
		WithQuery(`SELECT * FROM "users"  WHERE "users"."deleted_at" IS NULL AND (("users"."login" = test) AND ("users"."password" = test_password)) ORDER BY "users"."id" ASC LIMIT 1`).
		WithReply(commonReply)
	mocket.Catcher.Logging = true

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mockedRouter.ServeHTTP(w, req)

	assert.Equal(t, `{"token":"c1c17c916ba11acb38b41a3d99dc678a5b3a3d78"}`, w.Body.String())
}

func TestProfileUserUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.Setup()

	w := performRequest(r, "GET", "/api/user/profile")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, `{"error":"auth required"}`, w.Body.String())
}


func TestProfileUser(t *testing.T) {

	u := user.User{}
	u.ID = 123
	u.Login = "test"
	u.Password = "test_password"

	cache.InvalidateUserTokens()
	cache.AddUserToken(&u, "c1c17c916ba11acb38b41a3d99dc678a5b3a3d78")

	commonReply := []map[string]interface{}{{"id": "11", "media": nil}}

	mocket.Catcher.NewMock().
	WithQuery(`SELECT * FROM "playlists"  WHERE "playlists"."deleted_at" IS NULL AND (("user_id" = 123))`).
		WithReply(commonReply)

	req, _ := http.NewRequest("GET", "/api/user/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-AUTH-TOKEN", "c1c17c916ba11acb38b41a3d99dc678a5b3a3d78")

	w := httptest.NewRecorder()
	mockedRouter.ServeHTTP(w, req)

	receivedUser := &user.User{}

	json.Unmarshal(w.Body.Bytes(), receivedUser)

	assert.Equal(t, uint(123), receivedUser.ID)
	assert.Equal(t, uint(11), receivedUser.Playlists[0].ID)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)

	req.Header.Set("token", "test_token")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}


