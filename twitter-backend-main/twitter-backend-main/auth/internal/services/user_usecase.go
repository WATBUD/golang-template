package services

import (
	"github.com/leoantony72/twitter-backend/auth/internal/ports"
)

type userUseCase struct {
	userRepo UserUseCase22
}

type UserUseCase22 interface {

	//DoesKeyExist(u *UserPostgresRepo, key string)
}

func NewUseCase(repo UserUseCase22) ports.UserUseCase {
	return &userUseCase{userRepo: repo}
}

// func DoesKeyExist(u *UserPostgresRepo, key string) bool {

// }
