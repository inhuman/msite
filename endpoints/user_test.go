package endpoints_test

import (
	"testing"
	"net/http/httptest"
	"github.com/inhuman/msite/router"
	"net/http"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"bytes"
	"github.com/jinzhu/gorm"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"log"
	"github.com/inhuman/msite/db"
	"fmt"
	"time"
)

func TestRegisterUser(t *testing.T) {

	dbm, mock := getMock(t)
	defer dbm.Close()

	db.Stor.SetDb(dbm)

	tt := time.Now()

	//rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "login", "password"}).
	//	AddRow(1, tt, tt, nil, "test", "test_token")

	mock.ExpectExec(`INSERT INTO "users" (.+)`).
		WithArgs(tt, tt, nil, "test", "test_password").
		WillReturnResult(sqlmock.NewResult(1, 1))

	gin.SetMode(gin.TestMode)
	r := router.Setup()

	var jsonStr = []byte(`{
	"login": "test",
	"password": "test_password",
	"confirm_password": "test_password"
}`)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)


	fmt.Println(w.Body.String())


	endExpect(t, mock)

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

func getMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	dbm, mck, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, gerr := gorm.Open("postgres", dbm)
	if gerr != nil {
		log.Fatalf("can't open gorm connection: %s", err)
	}

	return gormDB.Set("gorm:update_column", true), mck
}

func endExpect(t *testing.T, mock sqlmock.Sqlmock) {
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
