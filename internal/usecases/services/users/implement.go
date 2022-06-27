package users

import (
	"context"
	"database/sql"

	"github.com/datshiro/crud/internal/infras/errors"
	"github.com/datshiro/crud/internal/infras/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	ctx  context.Context
	exec boil.ContextExecutor
}

var changesCols = boil.Whitelist(
	models.UserColumns.Name,
	models.UserColumns.Email,
)

func (u *userService) FindById(id int) (*models.User, error) {
	model, err := models.FindUser(u.ctx, u.exec, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.DataNotFoundError
		}
		return nil, err
	}
	if model == nil {
		return nil, errors.DataNotFoundError
	}
	return model, nil
}

func (u *userService) Create(name string, email string) (*models.User, error) {
	user := &models.User{Name: name, Email: null.StringFrom(email)}
	if err := user.Insert(u.ctx, u.exec, changesCols); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) FindByIdForUpdate(id int) (*models.User, error) {
	model, err := models.Users(models.UserWhere.ID.EQ(id),
		qm.For("update")).One(u.ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (u *userService) Update(user *models.User) error {
	_, err := user.Update(u.ctx, u.exec, changesCols)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) DeleteById(id int) error {
	_, err := models.Users(models.UserWhere.ID.EQ(id)).DeleteAll(u.ctx, u.exec, false)
	return err
}

