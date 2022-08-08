package storage

import (
	"errors"
	"os"
	"strconv"
)

const minSecretKeySize = 32

func GetSecretKey() ([]byte, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) < minSecretKeySize {
		return nil, errors.New("Must provide a secret key under env variable SECRET_KEY, length must be > " + strconv.Itoa(minSecretKeySize))
	}

	return []byte(secretKey), nil
}
