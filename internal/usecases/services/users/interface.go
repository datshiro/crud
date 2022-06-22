package users

import (
	"context"

	"github.com/datshiro/crud/internal/infras/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewService(ctx context.Context, exec boil.ContextExecutor) UserService {
	return &userService{
		ctx:  ctx,
		exec: exec,
	}
}

type UserService interface {
	Create(name string, email string) (*models.User, error)
}
