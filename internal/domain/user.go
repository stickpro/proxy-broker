package domain

type User struct {
	Id         uint64 `json:"id"`
	ProxyLogin string `json:"proxy_login"`
	ProxyPass  string `json:"proxy_pass"`
	Status     uint8  `json:"status"`
}

type UserKafkaMsg struct {
	Type string `json:"type"`
	Data User   `json:"data"`
}

func (u User) ToKafkaMsg() UserKafkaMsg {
	return UserKafkaMsg{
		Type: "credential",
		Data: u,
	}
}
