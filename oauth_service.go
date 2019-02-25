package oauth_service

import "golang.org/x/oauth2"

type Provider interface {
	Login(redirectURL, state string) error
	Config() oauth2.Config
}
