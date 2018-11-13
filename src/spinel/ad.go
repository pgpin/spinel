package spinel

import "gopkg.in/korylprince/go-ad-auth.v2"

type ADConnection struct {
	auth.Config
}

func (adc *ADConnection) Authenticate(user string, pass string) bool {
	status, err := auth.Authenticate(adc.config, user, pass)
	if err != nil || !status {
		return false
	} else {
		return true
	}
}
