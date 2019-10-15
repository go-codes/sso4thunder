package sso4thunder

import (
	"context"
	"golang.org/x/oauth2"
)


type SSO struct {
	conf Config
	oauthConf *oauth2.Config
}


func New (conf Config) *SSO {
	oauthConf := &oauth2.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Endpoint:     oauth2.Endpoint{
			AuthURL: conf.Endpoint + "v1/oauth2/authorize",
			TokenURL: conf.Endpoint + "v1/oauth2/token",
		},
		RedirectURL:  conf.RedirectUrl,
		Scopes:       conf.Scope,
	}
	return &SSO{
		conf: conf,
		oauthConf: oauthConf,
	}
}

func (s *SSO) AuthorizeURL (state string, opts ...oauth2.AuthCodeOption) string {
	return s.oauthConf.AuthCodeURL(state, opts...)
}

func (s *SSO) ExchangeToken (ctx context.Context, code string) (*oauth2.Token, error){
	token, err := s.oauthConf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
