package repo

import "dbingo/model"

type UserRepositoryI interface {
	GetByID(int64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	Create(*model.User) error
}
