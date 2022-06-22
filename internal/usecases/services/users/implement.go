package users

import (
	"context"

	"github.com/datshiro/crud/internal/infras/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type userService struct {
	ctx  context.Context
	exec boil.ContextExecutor
}

var changeColumns = boil.Whitelist(
	models.UserColumns.Name,
	models.UserColumns.Email,
)

func (u *userService) Create(name string, email string) (*models.User, error) {
	user := &models.User{Name: name, Email: null.StringFrom(email)}
	if err := user.Insert(u.ctx, u.exec, changeColumns); err != nil {
		return nil, err
	}
	return user, nil
}
