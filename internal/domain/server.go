package domain

import "time"

type Server struct {
	Id        uint64    `json:"id"`
	Ip        string    `json:"ip"`
	Status    uint8     `json:"status"`
	PortMin   uint16    `json:"port_min"`
	PortMax   uint16    `json:"port_max"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
