package repository

import (
	"asocks-ws/internal/domain"
	"asocks-ws/pkg/logger"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserProxy interface {
	FindByColumn(any, string) (domain.UserProxy, error)
}

func NewUserProxyRepository(db *gorm.DB) UserProxy {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) FindByColumn(value any, columnName string) (domain.UserProxy, error) {
	logger.Info("[UserProxyRepository]... Find by column")
	var userProxy domain.UserProxy
	err := u.DB.Find(&userProxy, columnName+" = ?", value).Error
	return userProxy, err
}
