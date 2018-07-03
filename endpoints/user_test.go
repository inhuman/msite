package endpoints_test

import (
	"testing"
	"net/http/httptest"
	"github.com/inhuman/msite/router"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"bytes"
	"github.com/inhuman/msite/db"
	"github.com/inhuman/msite/user"
	"encoding/json"
	mocket "github.com/Selvatico/go-mocket"
)

func TestRegisterUser(t *testing.T) {

	mocket.Catcher.Register()
	dbm, err := gorm.Open(mocket.DRIVER_NAME, "any_string")

	if err != nil {
		t.Fatal(err)
	}

	db.Stor.SetDb(dbm)

	gin.SetMode(gin.TestMode)
	r := router.Setup()

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
	r.ServeHTTP(w, req)


	u := &user.User{}

	json.Unmarshal(w.Body.Bytes(), u)

	assert.Equal(t, u.Login, "test")
	assert.Equal(t, u.Password, "********")

}


func TestLoginUser(t *testing.T) {

	mocket.Catcher.Register()
	dbm, err := gorm.Open(mocket.DRIVER_NAME, "any_string")

	if err != nil {
		t.Fatal(err)
	}

	db.Stor.SetDb(dbm)

	gin.SetMode(gin.TestMode)
	r := router.Setup()

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
	r.ServeHTTP(w, req)

	assert.Equal(t, `{"token":"c1c17c916ba11acb38b41a3d99dc678a5b3a3d78"}`, w.Body.String())

}

func TestProfileUserUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := router.Setup()

	w := performRequest(r, "GET", "/api/user/profile")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, `{"error":"auth required"}`, w.Body.String())
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)

	req.Header.Set("token", "test_token")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
