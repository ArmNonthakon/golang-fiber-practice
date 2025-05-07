package mapper

import (
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/generated/server"
)

type UserMapper interface {
	Mapper(res model.User) server.User
}

type UserMapperImpl struct {
}

func (u UserMapperImpl) Mapper(res model.User) server.User {
	return server.User{
		Id:   &res.ID,
		Name: res.Name,
	}
}

func NewUserMapper() UserMapper {
	return UserMapperImpl{}
}
