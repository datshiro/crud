package body

import (
	"fmt"
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/datshiro/crud/internal/infras/models"
	"github.com/stretchr/testify/assert"
)

var (
	mockName     = "dat_test"
	mockEmail    = "datshiro_test@gmail.com"
	mockBodyByte = []byte(fmt.Sprintf(`{"name":"%s","email":"%s"}`, mockName, mockEmail))
)

type BindRequestTestCase struct {
	BytesBody       []byte
	ExpectedErr     error
	ExpectedErrFunc func(*testing.T, *models.User)
}

var BindRequestTestCases = map[string]BindRequestTestCase{
	"should return nil error": {
		BytesBody: mockBodyByte,
		ExpectedErrFunc: func(t *testing.T, user *models.User) {
			assert.Equal(t, mockName, user.Name)
			assert.Equal(t, mockEmail, user.Email.String)
		},
		ExpectedErr: nil,
	},
	"should return invalid params error": {
		BytesBody:   []byte(""),
		ExpectedErr: errors.InvalidParamError,
	},
	"should return json error": {
		BytesBody:   []byte("{'name': 'dat', 'email': 'datshiro@gmail.com'}"),
		ExpectedErr: errors.InvalidParamError,
	},
}

func TestBindRequest(t *testing.T) {
	for tName, tCase := range BindRequestTestCases {
		t.Run(tName, func(t *testing.T) {
			user := &models.User{}
			assert.Equal(t, tCase.ExpectedErr, BindRequest(tCase.BytesBody, user))
			if tCase.ExpectedErrFunc != nil {
				tCase.ExpectedErrFunc(t, user)
			}
		})
	}
}
