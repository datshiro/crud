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
	"github.com/volatiletech/null/v8"
)

var (
	userCreatePayload = &models.User{Name: mockName, Email: null.StringFrom(mockEmail)}
)

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()

	h := NewPostHandler()
	e.POST("/users", h.Handle)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(mockUserJSON)+"\n"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := SetupContext(e, req, rec)

	// Assertions
	if assert.NoError(t, h.Handle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		u := &models.User{}
		err := body.BindRequest(rec.Body.Bytes(), u)
		assert.Nil(t, err)
		assert.Equal(t, userCreatePayload.Name, u.Name)
		assert.Equal(t, userCreatePayload.Email, u.Email)
	}
}
