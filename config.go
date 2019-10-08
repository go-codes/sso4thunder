package sso4thunder

type Config struct {
	Endpoint string
	ClientId string
	ClientSecret string
	RedirectUrl string
	Scope []string
}