package facebook

import "golang.org/x/oauth2"

type FacebookProvider struct {
	cfg oauth2.Config
}

func (f FacebookProvider) Login(redirectURL, state string) error {
	return nil
}

func (f FacebookProvider) Config() oauth2.Config {
	return f.cfg
}
