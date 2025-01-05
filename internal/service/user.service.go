package service

import (
	"github/manhwork/ecommerce_go/internal/repo"
)

type UserService struct {
	UserRepo *repo.Userrepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserrepo(),
	}
}

func (us *UserService) FindAllUsers() {
}

func (us *UserService) FindUserById(cccd string) {
}

func (us *UserService) Create() {
}
