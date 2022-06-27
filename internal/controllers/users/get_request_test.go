package users

import (
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/stretchr/testify/assert"
)

type GetRequestValidateTestCase struct {
	ID          int
	Page        int
	Limit       int
	ExpectedErr error
}

var GetRequestValidateTestCases = map[string]GetRequestValidateTestCase{
	"should return nil error": {
		ID:          1,
		ExpectedErr: nil,
	},
	"should return error on invalid id (negative)": {
		ID:          -1,
		Page:        1,
		Limit:       3,
		ExpectedErr: errors.InvalidIdError,
	},
	"should return error on invalid id": {
		ID:          0,
		Page:        0,
		Limit:       0,
		ExpectedErr: errors.InvalidParamError,
	},
	"should return error on page must be non negative": {
		ID:          0,
		Page:        -1,
		Limit:       5,
		ExpectedErr: errors.NewParamErr("page must be non negative"),
	},
	"should return error on limit must be non negative": {
		ID:          0,
		Page:        1,
		Limit:       -1,
		ExpectedErr: errors.NewParamErr("limit must be non negative"),
	},
}

func TestGetRequestValidate(t *testing.T) {
	for tName, tCase := range GetRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &getRequest{
				ID:    tCase.ID,
				Page:  tCase.Page,
				Limit: tCase.Limit,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}
