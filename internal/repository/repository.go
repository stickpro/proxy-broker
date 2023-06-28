package repository

import "gorm.io/gorm"

type Repositories struct {
	UserProxy UserProxy
	Server    Server
	User      User
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserProxy: NewUserProxyRepository(db),
		Server:    NewServerRepository(db),
		User:      NewUserRepository(db),
	}
}
