package users

import (
	"testing"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/stretchr/testify/assert"
)

type DeleteRequestValidateTestCase struct {
	ID          int
	ExpectedErr error
}

var DeleteRequestValidateTestCases = map[string]DeleteRequestValidateTestCase{
	"should return nil error": {
		ID:          1,
		ExpectedErr: nil,
	},
	"should return error on empty field": {
		ID:          0,
		ExpectedErr: errors.InvalidIdError,
	},
}

func TestDeleteRequestValidate(t *testing.T) {
	for tName, tCase := range DeleteRequestValidateTestCases {
		t.Run(tName, func(t *testing.T) {
			request := &deleteRequest{
				ID: tCase.ID,
			}
			assert.Equal(t, tCase.ExpectedErr, request.Validate())
		})
	}
}
