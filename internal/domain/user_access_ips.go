package domain

import (
	"gorm.io/gorm"
)

type UserAccessIps struct {
	Id        uint64         `json:"id"`
	UserId    uint64         `json:"user_id"`
	Ip        string         `json:"ip"`
	Status    int8           `json:"status"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
