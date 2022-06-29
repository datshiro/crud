package users

import (
	"testing"

	"github.com/datshiro/crud/internal/infras/models"
	"github.com/datshiro/crud/internal/usecases/services/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

var (
	mockID    = 1
	mockName  = "dat"
	mockEmail = "datshiro@gmail.com"
)

func TestServiceCreateUser(t *testing.T) {
	service := mocks.NewUserService(t)

	user := &models.User{Name: mockName, Email: null.StringFrom(mockEmail)}
	service.On("Create", mockName, mockEmail).Return(user, nil)

	createdUser, err := service.Create(mockName, mockEmail)
	assert.Equal(t, err, nil)
	assert.Equal(t, createdUser, user)
}

func TestServiceGetUser(t *testing.T) {
	service := mocks.NewUserService(t)

	user := &models.User{ID: mockID, Name: mockName, Email: null.StringFrom(mockEmail)}
	service.On("FindById", mockID).Return(user, nil)

	retrievedUser, err := service.FindById(mockID)
	assert.Equal(t, err, nil)
	assert.Equal(t, retrievedUser, user)
}
