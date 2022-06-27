package users

import (
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/stretchr/testify/assert"
)

type PostRequestValidateTestCase struct {
	Name        string
	Email       string
	ExpectedErr error
}

var PostRequestValidateTestCases = map[string]PostRequestValidateTestCase{
	"should return nil error": {
		Name:        "dat",
		Email:       "datshiro@gmail.com",
		ExpectedErr: nil,
	},
	"should return error on empty field": {
		Name:        "",
		Email:       "",
		ExpectedErr: errors.NewParamErr("Name must be provided"),
	},
}

func TestPostRequestValidate(t *testing.T) {
	for tName, tCase := range PostRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &postRequest{
				Name:  tCase.Name,
				Email: tCase.Email,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}
