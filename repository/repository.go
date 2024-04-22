package repository

import (
	"context"
	"database/sql"

	"github.com/orololuwa/crispy-octo-guacamole/models"
)

type UserRepo interface {
	CreateAUser(ctx context.Context, tx *sql.Tx, user models.User) (int, error)
	GetAUser(ctx context.Context, tx *sql.Tx, id int) (models.User, error)
	GetAllUser(ctx context.Context, tx *sql.Tx) ([]models.User, error)
	UpdateAUsersName(ctx context.Context, tx *sql.Tx, id int, firstName, lastName string)(error)
	DeleteUserByID(ctx context.Context, tx *sql.Tx, id int) error
}

type DBRepo interface {
	Transaction(ctx context.Context, operation func(context.Context, *sql.Tx) error) error 
}