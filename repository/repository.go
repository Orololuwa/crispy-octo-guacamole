package repository

import (
	"database/sql"

	"github.com/orololuwa/crispy-octo-guacamole/models"
)

type DatabaseRepo interface {
	CreateAUser(user models.User) (int, error)
	GetAUser(id int) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateAUsersName(id int, firstName, lastName string)(error)
	DeleteUserByID(id int) error
}

type Repo struct {
	DB *sql.DB
}

func New(conn *sql.DB) DatabaseRepo {
	return &Repo{
		DB: conn,
	}
}