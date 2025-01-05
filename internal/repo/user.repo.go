package repo

import (
	"github/manhwork/ecommerce_go/configs"
	"github/manhwork/ecommerce_go/internal/models"

	"gorm.io/gorm"
)

type Userrepo struct {
	DB *gorm.DB
}

func NewUserRepo() *Userrepo {
	db := configs.GetDBConfig()
	return &Userrepo{DB: db}
}

func (us *Userrepo) FindAllUsers() ([]models.User, error) {
	var users []models.User
	if err := us.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (us *Userrepo) FindUserById(cccd string) (*models.User, error) {
	var user models.User
	if err := us.DB.First(&user, "id = ?", cccd).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *Userrepo) Create(user *models.User) error {
	if err := us.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
