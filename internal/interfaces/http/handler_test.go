package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/domain/entity"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/generated/server"
	handler "github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/interfaces/http"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) DeleteUserId(id string) (string, error) {
	return "Delete success", nil
}

func (m *MockService) GetUser() ([]entity.UserEntity, error) {
	return []entity.UserEntity{{Id: "1", Name: "user1"}, {Id: "2", Name: "user2"}}, nil
}

func (m *MockService) GetUserId(id string) (entity.UserEntity, error) {
	if id == "1" {
		return entity.UserEntity{Id: "1", Name: "user1"}, nil
	} else if id == "2" {
		return entity.UserEntity{Id: "2", Name: "user2"}, nil
	}
	return entity.UserEntity{}, fmt.Errorf("Not found")
}

func (m *MockService) PostUser(name string) (entity.UserEntity, error) {
	return entity.UserEntity{Id: "3", Name: name}, nil
}

func (m *MockService) PutUserId(name string, id string) (entity.UserEntity, error) {
	return entity.UserEntity{Id: id, Name: name}, nil
}

func initService() server.ServerInterface {
	service := new(MockService)
	handler := handler.NewHandler(service)
	return handler
}

func TestGetUser(t *testing.T) {
	handler := initService()
	var users []model.User
	app := fiber.New()

	app.Get("/user", func(c *fiber.Ctx) error {
		return handler.GetUser(c)
	})

	req := httptest.NewRequest("GET", "/user", nil)
	resp, err := app.Test(req)

	err = json.NewDecoder(resp.Body).Decode(&users)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)

	assert.Equal(t, "1", users[0].ID)
	assert.Equal(t, "user1", *users[0].Name)
	assert.Equal(t, "2", users[1].ID)
	assert.Equal(t, "user2", *users[1].Name)

}

func TestGetUserById(t *testing.T) {
	handler := initService()
	var users model.User
	app := fiber.New()

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return handler.GetUserById(c, id)
	})

	req := httptest.NewRequest("GET", "/user/2", nil)
	resp, err := app.Test(req)

	err = json.NewDecoder(resp.Body).Decode(&users)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)

	assert.Equal(t, "2", users.ID)
	assert.Equal(t, "user2", *users.Name)

}

func TestDeleteUser(t *testing.T) {
	handler := initService()
	app := fiber.New()

	app.Delete("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return handler.DeleteUserById(c, id)
	})

	req := httptest.NewRequest("DELETE", "/user/1", nil)
	resp, err := app.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)

}

func TestCreateUser(t *testing.T) {
	handler := initService()
	var user model.User
	app := fiber.New()

	app.Post("/user", func(c *fiber.Ctx) error {
		return handler.CreateUser(c)
	})

	jsonBody := `{"name": "Arm"}`

	req := httptest.NewRequest("POST", "/user", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	err = json.NewDecoder(resp.Body).Decode(&user)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)

	assert.Equal(t, "3", user.ID)
	assert.Equal(t, "Arm", *user.Name)

}

func TestModifyUserById(t *testing.T) {
	handler := initService()
	var user model.User
	app := fiber.New()

	app.Put("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return handler.ModifyUserById(c, id)
	})

	jsonBody := `{"name": "New Name"}`

	req := httptest.NewRequest("PUT", "/user/1", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	err = json.NewDecoder(resp.Body).Decode(&user)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "New Name", *user.Name)

}
