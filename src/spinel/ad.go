package spinel

import "gopkg.in/korylprince/go-ad-auth.v2"

type ADConnection struct {
	Config *auth.Config
}

func NewActiveDirectoryConnection(server string, port int, basedn string) *ADConnection {
	config := &auth.Config{
		Server:   server,
		Port:     port,
		BaseDN:   basedn,
		Security: auth.SecurityNone,
	}

	return &ADConnection{Config: config}
}

func (adc *ADConnection) Authenticate(user string, pass string) bool {
	status, err := auth.Authenticate(adc.Config, user, pass)
	if err != nil || !status {
		return false
	} else {
		return true
	}
}
