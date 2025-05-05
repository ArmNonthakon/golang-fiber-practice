package mapper

import (
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
	domain "github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/domain/entity"
)

type UserMapper interface {
	Mapper(res model.User) domain.UserEntity
}

type UserMapperImpl struct {
}

func (u UserMapperImpl) Mapper(res model.User) domain.UserEntity {
	return domain.UserEntity{
		Id:   res.ID,
		Name: *res.Name,
	}
}

func NewUserMapper() UserMapper {
	return UserMapperImpl{}
}
