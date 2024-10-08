package db

import (
	"database/sql"
	"fmt"
	"gin-api-sample/internal/user/domain/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	_, err := r.db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	return err
}

func (r *userRepository) UpdateUser(user *model.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID)
	return err
}

func (r *userRepository) DeleteUser(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
