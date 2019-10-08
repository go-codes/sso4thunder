package sso4thunder

import
(
	"context"
	"encoding/json"
	"io/ioutil"
)

// SSO用户信息
type UserInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`

}

// 使用Code获取用户信息
func (s *SSO) UserInfo (ctx context.Context, code string) (*UserInfo, error){
	token, err := s.ExchangeToken(ctx, code)
	if err != nil {
		return nil, err
	}

	client := s.oauthConf.Client(ctx, token)
	resp, err := client.Get(s.conf.Endpoint + "v1/oauth2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user UserInfo
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}