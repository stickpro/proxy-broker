package repository

import (
	"asocks-ws/internal/domain"
	"asocks-ws/pkg/logger"
	"gorm.io/gorm"
)

type User interface {
	GetAll() ([]domain.User, error)
	GetByID(ID int) (domain.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return userRepository{
		DB: db,
	}
}

func (c userRepository) GetAll() ([]domain.User, error) {
	logger.Info("[User]...Get All")

	var res []domain.User
	if err := c.DB.Model(&domain.User{}).Find(&res).Error; err != nil {
		return nil, err
	}

	if err := c.DB.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (c userRepository) GetByID(ID int) (domain.User, error) {
	logger.Info("[User]...Get By ID")

	user := domain.User{}
	if err := c.DB.Where("id", ID).Find(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}
