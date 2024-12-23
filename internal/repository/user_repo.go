package repository

import (
	"database/sql"
	"user/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user domain.User) error {
	query := "INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Username, user.Password, user.CreatedAt)
	return err
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	query := "SELECT id, username, password, created_at FROM users WHERE username = $1"
	row := r.db.QueryRow(query, username)
	var user domain.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
