package users

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/datshiro/crud/internal/infras/db"
	"github.com/datshiro/crud/internal/infras/inject"
	"github.com/labstack/echo/v4"
)

func newMockUser(name, email string) []byte {
	return []byte(fmt.Sprintf(`{"name":"%s","email":"%s"}`, name, email))
}
func newMockUpdateUser(id int, name, email string) []byte {
	return []byte(fmt.Sprintf(`{"id": %d, "name":"%s","email":"%s"}`, id, name, email))
}

var (
	DbUrl        = "postgres://postgres:postgres@localhost:5432/crud?sslmode=disable"
	mockName     = "dat_test"
	mockEmail    = "datshiro_test@gmail.com"
	mockUserJSON = []byte(fmt.Sprintf(`{"name":"%s","email":"%s"}`, mockName, mockEmail))
	mockUpdateId = 3

	mockUpdatedName       = "updated_dat_test"
	mockUpdatedEmail      = "updated_datshiro_test@gmail.com"
	mockUserJSONForUpdate = []byte(fmt.Sprintf(`{"id":%d, "name":"%s","email":"%s"}`, mockUpdateId, mockUpdatedName, mockUpdatedEmail))

	mockDeleteUserID = "19"
)

func SetupContext(e *echo.Echo, req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	c := e.NewContext(req, rec)
	dbc, err := db.NewDB(DbUrl)
	if err != nil {
		fmt.Println(err)
		panic("Invalid database connection")
	}
	inject.SetDB(c, dbc)
	return c
}
