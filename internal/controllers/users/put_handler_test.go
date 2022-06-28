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

type UpdateUserTestCase struct {
	ID              int
	Name            string
	Email           string
	ExpectedErr     error
	ExpectedErrFunc func(*testing.T, *httptest.ResponseRecorder, error)
}

var UpdateUserTestCases = map[string]UpdateUserTestCase{
	"should return nil error": {
		ID:          mockUpdateId,
		Name:        mockUpdatedName,
		Email:       mockUpdatedEmail,
		ExpectedErr: nil,
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			object := &models.User{}
			_ = body.BindRequest(rec.Body.Bytes(), object)
			assert.Equal(t, mockUpdatedName, object.Name)
			assert.Equal(t, mockUpdatedEmail, object.Email.String)
		},
	},
	"should return invalid bind error": {
		ID:          -1,
		Name:        mockUpdatedName,
		Email:       mockUpdatedEmail,
		ExpectedErr: errors.InvalidIdError,
	},
	"should return data not found error": {
		ID:          999,
		Name:        mockUpdatedName,
		Email:       mockUpdatedEmail,
		ExpectedErr: errors.DataNotFoundError,
	},
}

func TestPutUser(t *testing.T) {
	// Setup
	e := echo.New()

	h := NewPutHandler()
	e.PUT("/users", h.Handle)
	for tName, tCase := range UpdateUserTestCases {
		t.Run(tName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(newMockUpdateUser(tCase.ID, tCase.Name, tCase.Email))+"\n"))
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
