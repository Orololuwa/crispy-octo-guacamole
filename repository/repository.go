package repository

import (
	"context"
	"database/sql"

	"github.com/orololuwa/crispy-octo-guacamole/models"
)

type UserRepo interface {
	CreateAUser(user models.User) (int, error)
	GetAUser(id int) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateAUsersName(id int, firstName, lastName string)(error)
	DeleteUserByID(id int) error
}

type DBRepo interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error 
}