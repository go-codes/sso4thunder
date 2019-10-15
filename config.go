package sso4thunder

type Config struct {
	Endpoint string `toml: "endpoint"`
	ClientId string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	RedirectUrl string `toml:"redirect_url"`
	Scope []string	`toml:"scope"`
}