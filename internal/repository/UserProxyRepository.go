package repository

import (
	"asocks-ws/internal/domain"
	"asocks-ws/pkg/logger"
	"gorm.io/gorm"
)

type userProxyRepository struct {
	DB *gorm.DB
}

type UserProxy interface {
	GetAll() ([]domain.UserProxy, error)
	FindByColumn(any, string) (domain.UserProxy, error)
	UpdateOne(domain.MessageUserProxy, domain.UpdateUserProxy) (domain.UserProxy, error)
}

func NewUserProxyRepository(db *gorm.DB) UserProxy {
	return userProxyRepository{
		DB: db,
	}
}

func (u userProxyRepository) FindByColumn(value any, columnName string) (domain.UserProxy, error) {
	logger.Info("[UserProxyRepository]... Find by column")
	var userProxy domain.UserProxy
	err := u.DB.Table("user_proxies").Find(&userProxy, columnName+" = ?", value).Error
	return userProxy, err
}

func (u userProxyRepository) GetAll() (userProxy []domain.UserProxy, err error) {
	logger.Info("[UserProxyRepository]...Get All")
	err = u.DB.Table("user_proxies").Find(&userProxy).Error
	return userProxy, err
}

func (u userProxyRepository) UpdateOne(message domain.MessageUserProxy, input domain.UpdateUserProxy) (domain.UserProxy, error) {
	var userProxy domain.UserProxy
	err := u.DB.Table("user_proxies").Model(&userProxy).Where("id = ?", message.ID).Update("extip", input.ExtIp).Error
	return userProxy, err
}
