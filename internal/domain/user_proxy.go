package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserProxy struct {
	id                   uint64
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
	DeletedAt            time.Time       `json:"deleted_at"`
	UserId               int             `json:"user_id"`
	ServerId             int             `json:"server_id"`
	Port                 int             `json:"port"`
	TypeId               int             `json:"type_id"`
	AuthTypeId           int             `json:"auth_type_id"`
	CountryCode          string          `json:"country_code"`
	StateId              int             `json:"state_id"`
	CityId               int             `json:"city_id"`
	Asn                  int             `json:"asn"`
	Status               int             `json:"status"`
	Password             string          `json:"password"`
	ExtIp                string          `json:"ext_ip"`
	ExtIpStatus          string          `json:"ext_ip_status"`
	Session              int             `json:"session"`
	MethodRotate         string          `json:"method_rotate"`
	Timeout              int             `json:"timeout"`
	ProxyTypeId          int8            `json:"proxy_type_id"`
	Name                 string          `json:"name"`
	LastConnectedAt      time.Time       `json:"last_connected_at"`
	CostLastConnectedDay decimal.Decimal `json:"cost_last_connected_day"`
}
