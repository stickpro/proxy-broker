package repository

import (
	"asocks-ws/internal/domain"
	"asocks-ws/pkg/logger"
	"gorm.io/gorm"
)

type serverRepository struct {
	DB *gorm.DB
}

type Server interface {
	FindByColumn(any, string) (domain.Server, error)
}

func NewServerRepository(DB *gorm.DB) Server {
	return &serverRepository{DB: DB}
}

func (u serverRepository) FindByColumn(value any, columnName string) (domain.Server, error) {
	logger.Info("[ServerRepository]... Find by column")
	var server domain.Server
	err := u.DB.Table("servers").Find(&server, columnName+" = ?", value).Error
	return server, err
}
