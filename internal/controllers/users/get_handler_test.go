package users

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type GetUserTestCase struct {
	ID              string
	Limit           string
	Page            string
	ExpectedErr     error
	ExpectedErrFunc func(*testing.T, *httptest.ResponseRecorder, error)
}

var GetUserTestCases = map[string]GetUserTestCase{
	"should return nil error": {
		ID: "2",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, userJSON, rec.Body.String())
		},
	},
	"should return not found error": {
		ID: "100",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
		},
		ExpectedErr: errors.DataNotFoundError,
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
var GetPaginationUserTestCases = map[string]GetUserTestCase{
	"should return nil error": {
		Page:  "1",
		Limit: "2",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEmpty(t, rec.Body.String())
		},
	},
	"should return invalid error": {
		Page:  "0",
		Limit: "0",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
		},
		ExpectedErr: errors.InvalidParamError,
	},
	"should return page non negative error": {
		Page:  "-1",
		Limit: "2",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
		},
		ExpectedErr: errors.NewParamErr("page must be non negative"),
	},
	"should return limit non negative error": {
		Page:  "1",
		Limit: "-2",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.Equal(t, http.StatusOK, rec.Code)
		},
		ExpectedErr: errors.NewParamErr("limit must be non negative"),
	},
	"should return invalid bind error": {
		Page:  "abc",
		Limit: "xyz",
		ExpectedErrFunc: func(t *testing.T, rec *httptest.ResponseRecorder, err error) {
			assert.IsType(t, &echo.HTTPError{}, err)
		},
	},
}

var (
	userJSON = `{"id":2,"name":"dat_test","email":"datshiro_test@gmail.com"}` + "\n"
)

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewGetHandler()
	for tName, tCase := range GetUserTestCases {
		t.Run(tName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
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
			tCase.ExpectedErrFunc(t, rec, err)
		})
	}
}

func TestGetPaginationUser(t *testing.T) {
	// Setup
	e := echo.New()
	h := NewGetHandler()
	for tName, tCase := range GetPaginationUserTestCases {
		t.Run(tName, func(t *testing.T) {
			// Build request
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			q := req.URL.Query()
			q.Add("limit", tCase.Limit)
			q.Add("page", tCase.Page)
			req.URL.RawQuery = q.Encode()

			rec := httptest.NewRecorder()
			c := SetupContext(e, req, rec)
			c.SetPath("/users")

			err := h.Handle(c)
			if tCase.ExpectedErr != nil {
				assert.Equal(t, tCase.ExpectedErr, err)
			}
			// Assertions
			tCase.ExpectedErrFunc(t, rec, err)
		})
	}
}
