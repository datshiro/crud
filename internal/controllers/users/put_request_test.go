package users

import (
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/stretchr/testify/assert"
)

type PutRequestValidateTestCase struct {
	Name        string
	Email       string
	ID          int
	ExpectedErr error
}

var PutRequestValidateTestCases = map[string]PutRequestValidateTestCase{
	"should return nil error": {
		ID:          1,
		Name:        "dat",
		Email:       "datshiro@gmail.com",
		ExpectedErr: nil,
	},
	"should return invalid id error": {
		ID:          0,
		Name:        "dat",
		Email:       "datshiro@gmail.com",
		ExpectedErr: errors.InvalidIdError,
	},
	"should return error on empty field": {
		ID:          2,
		Name:        "",
		Email:       "",
		ExpectedErr: errors.NewParamErr("Name must be provided"),
	},
}

func TestPutRequestValidate(t *testing.T) {
	for tName, tCase := range PutRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &putRequest{
				Name:  tCase.Name,
				Email: tCase.Email,
				ID:    tCase.ID,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}
