package main

import (
	"log"

	db "github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database"
	data "github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/repository"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/generated/server"
	handler "github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/interfaces/http"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/pkg/mapper"
	"github.com/joho/godotenv"

	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/domain/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	newDb := db.NewDb()
	repo := data.NewRepository(newDb)
	mapper := mapper.NewUserMapper()
	usecase := usecase.NewService(repo, mapper)
	handler := handler.NewHandler(usecase)

	server.RegisterHandlers(app, handler)

	app.Listen(":3000")
}
