package users

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/datshiro/crud/internal/infras/models"
	"github.com/datshiro/crud/internal/utils/body"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPutUser(t *testing.T) {
	// Setup
	e := echo.New()

	h := NewPutHandler()
	e.POST("/users", h.Handle)
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(mockUserJSONForUpdate)+"\n"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := SetupContext(e, req, rec)

	// Assertions
	if assert.NoError(t, h.Handle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		u := &models.User{}
		err := body.BindRequest(rec.Body.Bytes(), u)
		assert.Nil(t, err)
		assert.Equal(t, mockUpdatedName, u.Name)
		assert.Equal(t, mockUpdatedEmail, u.Email.String)
	}
}
