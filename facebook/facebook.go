package facebook

import "golang.org/x/oauth2"

type FacebookProvider struct {
	cfg oauth2.Config
}

var Provider FacebookProvider

func (f FacebookProvider) Login(redirectURL, state string) error {
	return nil
}

func (f FacebookProvider) Config() oauth2.Config {
	return f.cfg
}

func (f FacebookProvider) SetClientID(client_id string) {
	f.cfg.ClientID = client_id
}

func (f FacebookProvider) SetClientSecret(client_secret string) {
	f.cfg.ClientSecret = client_secret
}
