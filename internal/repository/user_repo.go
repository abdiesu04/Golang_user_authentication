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

func (r *UserRepository) SaveRefreshToken(userID int64, refreshToken string) error {
	query := "INSERT INTO refresh_tokens (user_id, token) VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET token = $2"
	_, err := r.db.Exec(query, userID, refreshToken)
	return err
}

func (r *UserRepository) GetRefreshToken(userID int64) (string, error) {
	query := "SELECT token FROM refresh_tokens WHERE user_id = $1"
	row := r.db.QueryRow(query, userID)
	var token string
	err := row.Scan(&token)
	return token, err
}

func (r *UserRepository) InvalidateRefreshToken(userID int64) error {
	query := "DELETE FROM refresh_tokens WHERE user_id = $1"
	_, err := r.db.Exec(query, userID)
	return err
}
