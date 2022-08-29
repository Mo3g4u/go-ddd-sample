package repository

import "github.com/Mo3g4u/go-ddd-sample/domain/model"

type Repository interface {
	FindUserByName(name string) (*model.User, error)
	CreateUser(user *model.User) error
}
