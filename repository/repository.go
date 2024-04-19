package repository

import (
	"github.com/orololuwa/crispy-octo-guacamole/models"
)

type UserRepo interface {
	CreateAUser(user models.User) (int, error)
	GetAUser(id int) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateAUsersName(id int, firstName, lastName string)(error)
	DeleteUserByID(id int) error
}

