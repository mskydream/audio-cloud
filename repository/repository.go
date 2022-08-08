package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mskydream/audio-cloud/storage"
)

type Authorization interface {
	CreateUser(user storage.User) (int, error)
	GetUser(username, password string) (storage.User, error)
	SetRefreshToken(username int, refreshToken string, refreshTokenTTL time.Duration) error
	UpdateRefreshToken(oldRefreshToken string, newRefreshToken string, refreshTokenTTL time.Duration) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB, dirPath string) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
