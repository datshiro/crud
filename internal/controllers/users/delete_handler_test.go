package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type DeleteUserTestCase struct {
	ID              string
	ExpectedErr     error
	ExpectedErrFunc func(*testing.T, *httptest.ResponseRecorder, error)
}

var DeleteUserTestCases = map[string]DeleteUserTestCase{
	"should return nil error": {
		ID: mockDeleteUserID,
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "null\n", rec.Body.String())
		},
	},
	"should return not found error": {
		ID: "999",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
		},
	},
	"should return invalid id error": {
		ID: "-1",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
		},
		ExpectedErr: errors.InvalidIdError,
	},
	"should return invalid bind error": {
		ID: "abc",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.IsType(t, &echo.HTTPError{}, err)
		},
	},
}

func TestDeleteUser(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewDeleteHandler()
	for tName, tCase := range DeleteUserTestCases {
		t.Run(tName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/", nil)
			rec := httptest.NewRecorder()
			c := SetupContext(e, req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(tCase.ID)

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
