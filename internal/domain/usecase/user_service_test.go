package usecase_test

import (
	"fmt"
	"testing"

	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/domain/usecase"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/generated/server"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/pkg/mapper"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/pkg/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) DeleteUserId(id string) (string, error) {
	if id == "1" || id == "2" {
		return "Delete success", nil
	}
	return "", fmt.Errorf("Not found")
}

func (m *MockRepository) GetUser() ([]model.User, error) {
	return []model.User{{ID: "1", Name: pointer.Ptr("user1")}, {ID: "2", Name: pointer.Ptr("user2")}}, nil
}

func (m *MockRepository) GetUserId(id string) (model.User, error) {
	if id == "1" {
		return model.User{ID: "1", Name: pointer.Ptr("user1")}, nil
	} else if id == "2" {
		return model.User{ID: "2", Name: pointer.Ptr("user2")}, nil
	}
	return model.User{}, fmt.Errorf("Not found")

}

func (m *MockRepository) PostUser(name string) (model.User, error) {
	return model.User{ID: "3", Name: pointer.Ptr(name)}, nil
}

func (m *MockRepository) PutUserId(name string, id string) (model.User, error) {
	if id == "1" {
		return model.User{ID: "1", Name: pointer.Ptr(name)}, nil
	} else if id == "2" {
		return model.User{ID: "2", Name: pointer.Ptr(name)}, nil
	}
	return model.User{}, fmt.Errorf("Not found")

}

func initService() (usecase.UserService, *MockRepository) {
	repo := new(MockRepository)
	mapper := mapper.UserMapperImpl{}
	service := usecase.NewService(repo, mapper)
	return service, repo
}

func TestGetUser(t *testing.T) {
	service, mockRepo := initService()
	expectedUsers := []server.User{{Id: pointer.Ptr("1"), Name: pointer.Ptr("user1")}, {Id: pointer.Ptr("2"), Name: pointer.Ptr("user2")}}

	actual, err := service.GetUser()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockRepo.AssertExpectations(t)
}

func TestGetUserId(t *testing.T) {
	service, mockRepo := initService()
	expectedUsers1 := server.User{Id: pointer.Ptr("1"), Name: pointer.Ptr("user1")}
	expectedUsers2 := server.User{Id: pointer.Ptr("2"), Name: pointer.Ptr("user2")}

	actual1, err := service.GetUserId("1")
	actual2, err := service.GetUserId("2")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers1, actual1)
	mockRepo.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers2, actual2)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	service, mockRepo := initService()
	expectedUsers := "Delete success"

	actual, err := service.DeleteUserId("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockRepo.AssertExpectations(t)
}

func TestPostUser(t *testing.T) {
	service, mockRepo := initService()
	expectedUsers := server.User{Id: pointer.Ptr("3"), Name: pointer.Ptr("newUser")}

	actual, err := service.PostUser("newUser")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockRepo.AssertExpectations(t)
}

func TestPutUserId(t *testing.T) {
	service, mockRepo := initService()
	expectedUsers := server.User{Id: pointer.Ptr("1"), Name: pointer.Ptr("newName")}

	actual, err := service.PutUserId("newName", "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockRepo.AssertExpectations(t)
}
