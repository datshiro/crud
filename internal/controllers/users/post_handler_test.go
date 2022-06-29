package users

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/datshiro/crud/internal/infras/models"
	"github.com/datshiro/crud/internal/utils/body"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type CreateUserTestCase struct {
	Name            string
	Email           string
	ExpectedErr     error
	ExpectedErrFunc func(*testing.T, *httptest.ResponseRecorder, error)
}

var CreateUserTestCases = map[string]CreateUserTestCase{
	"should return nil error": {
		Name:        mockName,
		Email:       mockEmail,
		ExpectedErr: nil,
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			object := &models.User{}
			_ = body.BindRequest(rec.Body.Bytes(), object)
			assert.Equal(t, mockName, object.Name)
			assert.Equal(t, mockEmail, object.Email.String)
		},
	},
	"should return invalid bind error": {
		Name:        "",
		Email:       "",
		ExpectedErr: errors.NewParamErr("Name must be provided"),

		// ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
		// 	assert.IsType(t, &echo.HTTPError{}, err)
		// },
	},
}

func TestCreateUser(t *testing.T) {
	// Setup
	e := echo.New()

	h := NewPostHandler()
	for tName, tCase := range CreateUserTestCases {
		t.Run(tName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(newMockUser(tCase.Name, tCase.Email))+"\n"))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := SetupContext(e, req, rec)

			// Assertions
			err := h.Handle(c)
			if tCase.ExpectedErr != nil {
				assert.Equal(t, tCase.ExpectedErr, err)
			}

			// Assertions
			if tCase.ExpectedErrFunc != nil {
				tCase.ExpectedErrFunc(t, rec, err)
			}
		})
	}
}
