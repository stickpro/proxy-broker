package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserProxy struct {
	Id                   uint64          `json:"id"`
	IP                   string          `json:"ip"`
	UserId               uint64          `json:"user_id"`
	ServerId             int             `json:"server_id"`
	Port                 uint32          `json:"port"`
	TypeId               int             `json:"type_id"`
	AuthTypeId           int8            `json:"auth_type_id"`
	CountryCode          string          `json:"country_code"`
	StateId              int             `json:"state_id,omitempty"`
	CityId               int             `json:"city_id,omitempty"`
	Asn                  int             `json:"asn,omitempty"`
	Status               int             `json:"status"`
	Password             string          `json:"password"`
	ExtIp                string          `json:"extip,omitempty"`
	ExtIpStatus          string          `json:"ext_ip_status"`
	Session              uint64          `json:"session"`
	MethodRotate         string          `json:"method_rotate"`
	Timeout              int             `json:"timeout"`
	ProxyTypeId          int8            `json:"proxy_type_id"`
	Name                 string          `json:"name"`
	CostLastConnectedDay decimal.Decimal `json:"cost_last_connected_day"`
	Login                string          `json:"login"`
	Ips                  []UserAccessIps `gorm:"foreignKey:UserId;references:UserId"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}

type UpdateUserProxy struct {
	ExtIp string `json:"extip"`
}

type MessageUserProxy struct {
	ID int    `json:"id"`
	IP string `json:"ip"`
}

type UserProxyForLpm struct {
	Id         uint64    `json:"id"`
	UserId     uint64    `json:"user_id"`
	LoginData  LoginData `json:"loginData"`
	ServerPort uint32    `json:"server_port"`
	Updated    int64     `json:"updated"`
	Auth       AuthProxy `json:"auth"`
	Status     int       `json:"status"`
}

type LoginData struct {
	Id      uint64 `json:"id"`
	Zone    string `json:"zone"`
	Country string `json:"country"`
	State   int    `json:"state,omitempty"`
	City    int    `json:"city,omitempty"`
	Asn     int    `json:"asn,omitempty"`
	LastIp  string `json:"lastip,omitempty"`
	Session uint64 `json:"session,omitempty"`
	Hold    string `json:"hold,omitempty"`
	Timeout int    `json:"timeout,omitempty"`
	Level   int32  `json:"level"`
}

type AuthProxy struct {
	Type     string   `json:"type"`
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Ips      []string `json:"ips"`
}
