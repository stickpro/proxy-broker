package enums

type ProxyType int8

const (
	all    ProxyType = 0
	res    ProxyType = 1
	any    ProxyType = 2
	mobile ProxyType = 3
	corp   ProxyType = 4
)

func (p ProxyType) GetProxyType(index int8) string {
	return []string{"all", "res", "any", "mobile", "corp"}[index]
}

type ProxyAuthType int8

const (
	password ProxyAuthType = iota + 1
	whitelistOrPassword
	whitelistAndPassword
	whitelist
)

func (p ProxyAuthType) GetProxyAuthType(index int8) string {
	// todo change normal min max struct
	if index == 0 {
		index = 1
	}
	return []string{"password", "whitelist_or_password", "whitelist_and_password", "whitelist"}[index-1]
}

const (
	KeepProxy              int = 1
	KeepConnection         int = 2
	RotateConnection       int = 3
	KeepConnectionLowTrust int = 4
)

const (
	MethodRotateManually    = "manually"
	MethodRotateEverRequest = "ever_request"
	MethodRotateTimeout     = "timeout"
)
