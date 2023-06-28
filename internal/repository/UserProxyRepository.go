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
	FindByColumn(any, string) ([]domain.UserProxy, error)
	FindByIp(ip string) ([]domain.UserProxy, error)
	UpdateOne(domain.MessageUserProxy, domain.UpdateUserProxy) (domain.UserProxy, error)
	LoadById(any) (domain.UserProxy, error)
}

func NewUserProxyRepository(db *gorm.DB) UserProxy {
	return userProxyRepository{
		DB: db,
	}
}

func (u userProxyRepository) FindByColumn(value any, columnName string) ([]domain.UserProxy, error) {
	logger.Info("[UserProxyRepository]... Find by column")
	var userProxy []domain.UserProxy
	err := u.DB.Model(&domain.UserProxy{}).Preload("Ips", "status = 1").Find(&userProxy, columnName+" = ?", value).Error
	return userProxy, err
}

func (u userProxyRepository) FindByIp(ip string) ([]domain.UserProxy, error) {
	logger.Info("[UserProxyRepository]... Find by Ip")
	var userProxy []domain.UserProxy
	err := u.DB.Joins("left join servers on servers.id = user_proxies.server_id").
		Where("servers.ip = ?", ip).
		Select("user_proxies.*").
		Preload("Ips", "status = 1").
		Find(&userProxy).Error
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

func (u userProxyRepository) LoadById(value any) (domain.UserProxy, error) {
	logger.Info("[UserProxyRepository]... Load By Id")
	var userProxy domain.UserProxy
	err := u.DB.Model(&domain.UserProxy{}).Preload("Ips", "status = 1").First(&userProxy, "id = ?", value).Error
	return userProxy, err
}
