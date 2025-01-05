package service

import (
	"github/manhwork/ecommerce_go/internal/models"
	"github/manhwork/ecommerce_go/internal/repo"
)

type UserService struct {
	UserRepo *repo.Userrepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) FindAllUsers() ([]models.User, error) {
	return us.UserRepo.FindAllUsers()
}

func (us *UserService) FindUserById(cccd string) (*models.User, error) {
	return us.UserRepo.FindUserById(cccd)
}

func (us *UserService) Create(user *models.User) error {
	return us.UserRepo.Create(user)
}
