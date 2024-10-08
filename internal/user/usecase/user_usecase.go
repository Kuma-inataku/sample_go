package usecase

import (
	"gin-api-sample/internal/user/domain/model"
	"gin-api-sample/internal/user/infrastructure/db"
)

type UserUsecase interface {
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
}

type userUsecase struct {
	userRepository db.UserRepository
}

func NewUserUsecase() UserUsecase {
	return &userUsecase{
		userRepository: db.NewUserRepository(),
	}
}

func (u *userUsecase) GetUsers() ([]model.User, error) {
	return u.userRepository.GetUsers()
}

func (u *userUsecase) CreateUser(user *model.User) error {
	return u.userRepository.CreateUser(user)
}

func (u *userUsecase) UpdateUser(user *model.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.userRepository.DeleteUser(id)
}
