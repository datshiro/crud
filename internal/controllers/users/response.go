package users

import (
	"github.com/datshiro/crud/internal/infras/models"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewResponse(model *models.User) *User {
	return &User{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email.String,
	}
}

func NewResponses(models models.UserSlice) []*User {
	resps := []*User{}
	for _, u := range models {
		resps = append(resps, NewResponse(u))
	}
	return resps
}
