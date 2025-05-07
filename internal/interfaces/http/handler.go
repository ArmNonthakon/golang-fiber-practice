package handler

import (
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
	usecase "github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/domain/usecase"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/generated/server"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase usecase.UserService
}

func (s *Handler) DeleteUserById(c *fiber.Ctx, id string) error {
	users, err := s.usecase.DeleteUserId(id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(users)
}

func (s *Handler) GetUser(c *fiber.Ctx) error {
	users, err := s.usecase.GetUser()
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return c.Status(fiber.StatusAccepted).JSON([]model.User{})
	}
	return c.Status(fiber.StatusAccepted).JSON(users)
}

func (s *Handler) GetUserById(c *fiber.Ctx, id string) error {
	user, err := s.usecase.GetUserId(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}
	return c.Status(fiber.StatusAccepted).JSON(user)
}

func (s *Handler) CreateUser(c *fiber.Ctx) error {
	user := new(server.CreateUserJSONRequestBody)

	if err := c.BodyParser(user); err != nil {
		return err
	}
	users, err := s.usecase.PostUser(*user.Name)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(users)
}

func (s *Handler) ModifyUserById(c *fiber.Ctx, id string) error {
	user := new(server.ModifyUserByIdJSONRequestBody)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	users, err := s.usecase.PutUserId(*user.Name, id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(users)
}

func NewHandler(usecase usecase.UserService) server.ServerInterface {
	return &Handler{usecase: usecase}
}
