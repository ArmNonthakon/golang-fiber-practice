package repository

import (
	db "github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
)

type UserRepository interface {
	GetUser() ([]model.User, error)
	PostUser(name string) (model.User, error)
	DeleteUserId(id string) (string, error)
	GetUserId(id string) (model.User, error)
	PutUserId(name string, id string) (model.User, error)
}

type UserRepositoryImpl struct {
	db db.DB
}

func (r *UserRepositoryImpl) DeleteUserId(id string) (string, error) {
	res, err := r.db.DeleteUserId(id)
	if err != nil {
		return "", err
	}
	return res, err
}

func (r *UserRepositoryImpl) GetUser() ([]model.User, error) {
	user, err := r.db.GetUser()
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserRepositoryImpl) GetUserId(id string) (model.User, error) {
	users, err := r.db.GetUserId(id)
	if err != nil {
		return model.User{}, err
	}
	return users, err
}

func (r *UserRepositoryImpl) PostUser(name string) (model.User, error) {
	user, err := r.db.PostUser(name)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func (r *UserRepositoryImpl) PutUserId(name string, id string) (model.User, error) {
	user, err := r.db.PutUserId(name, id)
	if err != nil {
		return model.User{}, err
	}
	return user, err
}

func NewRepository(db db.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}
