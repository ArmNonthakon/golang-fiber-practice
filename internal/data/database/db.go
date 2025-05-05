package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/model"
	"github.com/ArmNonthakon/golang-openapi-openapicodegen/internal/data/database/jet_generated/go_database/table"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type DB interface {
	GetUser() ([]model.User, error)
	PostUser(name string) (model.User, error)
	DeleteUserId(id string) (string, error)
	GetUserId(id string) (model.User, error)
	PutUserId(name string, id string) (model.User, error)
}

type DbImpl struct {
	SqlDb *sql.DB
}

func (db *DbImpl) GetUser() ([]model.User, error) {
	stmt := mysql.SELECT(table.User.ID, table.User.Name).FROM(table.User)

	var users []model.User
	err := stmt.Query(db.SqlDb, &users)
	return users, err
}

func (db *DbImpl) GetUserId(id string) (model.User, error) {
	stmt := mysql.
		SELECT(table.User.ID, table.User.Name).
		FROM(table.User).
		WHERE(table.User.ID.EQ(mysql.String(id))).
		LIMIT(1)
	var user model.User
	err := stmt.Query(db.SqlDb, &user)

	return user, err
}

func (db *DbImpl) PostUser(name string) (model.User, error) {
	id := uuid.New().String()

	stmt := table.User.INSERT(table.User.ID, table.User.Name).VALUES(id, name)

	_, err := stmt.Exec(db.SqlDb)

	return model.User{ID: id, Name: &name}, err
}

func (db *DbImpl) PutUserId(name string, id string) (model.User, error) {
	stmt := table.User.UPDATE(table.User.Name).SET(name).WHERE(table.User.ID.EQ(mysql.String(id)))

	_, err := stmt.Exec(db.SqlDb)
	return model.User{ID: id, Name: &name}, err
}

func (db *DbImpl) DeleteUserId(id string) (string, error) {
	stmt := table.User.DELETE().WHERE(table.User.ID.EQ(mysql.String(id)))

	_, err := stmt.Exec(db.SqlDb)
	return "Delete user successfully", err
}

func NewDb() DB {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name

	sqlDb, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error opening database: ", err)
		return &DbImpl{}
	}

	return &DbImpl{SqlDb: sqlDb}
}
