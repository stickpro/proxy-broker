package repository

import "gorm.io/gorm"

type Repositories struct {
	UserProxy UserProxy
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserProxy: NewUserProxyRepository(db),
	}
}
