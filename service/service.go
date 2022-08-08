package service

import (
	"time"

	"github.com/mskydream/audio-cloud/repository"
	"github.com/mskydream/audio-cloud/storage"
)

type Authorization interface {
	CreateUser(user storage.User) (int, error)
	GenerateAccessToken(username, password string) (int, string, error)
	UpdateAccessToken(userId int) (string, error)
	GenerateRefreshToken(userId int) (string, error)
	UpdateRefreshToken(oldRefreshToken string) (int, string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository, accessTokenTTL, refreshTokenTTL time.Duration) *Service {
	return &Service{
		Authorization: NewAuthService(repos, accessTokenTTL, refreshTokenTTL),
	}
}
