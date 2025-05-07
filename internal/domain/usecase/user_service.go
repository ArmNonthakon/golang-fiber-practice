package usecase

import (
	"fmt"

	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/repository"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/generated/server"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/pkg/mapper"
)

type UserService interface {
	GetUser() ([]server.User, error)
	PostUser(name string) (server.User, error)
	DeleteUserId(id string) (string, error)
	GetUserId(id string) (server.User, error)
	PutUserId(name string, id string) (server.User, error)
}

type UserServiceImpl struct {
	repo   repository.UserRepository
	mapper mapper.UserMapper
}

func (u *UserServiceImpl) GetUser() ([]server.User, error) {
	res, err := u.repo.GetUser()
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no data")
	}

	result := make([]server.User, 0, len(res))
	for _, data := range res {
		result = append(result, u.mapper.Mapper(data))
	}

	return result, nil
}

func (u *UserServiceImpl) GetUserId(id string) (server.User, error) {
	res, err := u.repo.GetUserId(id)

	if err != nil {
		return server.User{}, err
	}
	result := u.mapper.Mapper(res)

	if result.Id == nil {
		return server.User{}, fmt.Errorf("No data")
	}
	return result, nil
}

func (u *UserServiceImpl) DeleteUserId(id string) (string, error) {
	res, err := u.repo.DeleteUserId(id)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (u *UserServiceImpl) PostUser(name string) (server.User, error) {
	res, err := u.repo.PostUser(name)
	if err != nil {
		return server.User{}, err
	}
	result := u.mapper.Mapper(res)

	if result.Id == nil {
		return server.User{}, fmt.Errorf("No data")
	}
	return result, nil
}

func (u *UserServiceImpl) PutUserId(name string, id string) (server.User, error) {
	res, err := u.repo.PutUserId(name, id)
	if err != nil {
		return server.User{}, err
	}
	result := u.mapper.Mapper(res)

	if result.Id == nil {
		return server.User{}, fmt.Errorf("No data")
	}
	return result, nil
}

func NewService(repo repository.UserRepository, mapper mapper.UserMapper) UserService {
	return &UserServiceImpl{repo: repo, mapper: mapper}
}
