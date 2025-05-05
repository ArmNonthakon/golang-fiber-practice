package repository_test

import (
	"fmt"
	"testing"

	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/repository"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/pkg/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) DeleteUserId(id string) (string, error) {
	if id == "1" || id == "2" {
		return "Delete success", nil
	}
	return "", fmt.Errorf("Not found")
}

func (m *MockDB) GetUser() ([]model.User, error) {
	return []model.User{{ID: "1", Name: pointer.Ptr("user1")}, {ID: "2", Name: pointer.Ptr("user2")}}, nil
}

func (m *MockDB) GetUserId(id string) (model.User, error) {
	if id == "1" {
		return model.User{ID: "1", Name: pointer.Ptr("user1")}, nil
	} else if id == "2" {
		return model.User{ID: "2", Name: pointer.Ptr("user2")}, nil
	}
	return model.User{}, fmt.Errorf("Not found")

}

func (m *MockDB) PostUser(name string) (model.User, error) {
	return model.User{ID: "3", Name: pointer.Ptr(name)}, nil
}

func (m *MockDB) PutUserId(name string, id string) (model.User, error) {
	if id == "1" {
		return model.User{ID: "1", Name: pointer.Ptr(name)}, nil
	} else if id == "2" {
		return model.User{ID: "2", Name: pointer.Ptr(name)}, nil
	}
	return model.User{}, fmt.Errorf("Not found")

}

func initRepository() (repository.UserRepository, *MockDB) {
	db := new(MockDB)
	repo := repository.NewRepository(db)
	return repo, db
}

func TestGetUser(t *testing.T) {
	repo, mockDB := initRepository()
	expectedUsers := []model.User{{ID: "1", Name: pointer.Ptr("user1")}, {ID: "2", Name: pointer.Ptr("user2")}}

	actual, err := repo.GetUser()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockDB.AssertExpectations(t)
}

func TestGetUserId(t *testing.T) {
	repo, mockDB := initRepository()
	expectedUsers1 := model.User{ID: "1", Name: pointer.Ptr("user1")}
	expectedUsers2 := model.User{ID: "2", Name: pointer.Ptr("user2")}

	actual1, err := repo.GetUserId("1")
	actual2, err := repo.GetUserId("2")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers1, actual1)
	mockDB.AssertExpectations(t)

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers2, actual2)
	mockDB.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	repo, mockDB := initRepository()
	expectedUsers := "Delete success"

	actual, err := repo.DeleteUserId("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockDB.AssertExpectations(t)
}

func TestPostUser(t *testing.T) {
	repo, mockDB := initRepository()
	expectedUsers := model.User{ID: "3", Name: pointer.Ptr("newUser")}

	actual, err := repo.PostUser("newUser")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockDB.AssertExpectations(t)
}

func TestPutUserId(t *testing.T) {
	repo, mockDB := initRepository()
	expectedUsers := model.User{ID: "1", Name: pointer.Ptr("newName")}

	actual, err := repo.PutUserId("newName", "1")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actual)
	mockDB.AssertExpectations(t)
}
