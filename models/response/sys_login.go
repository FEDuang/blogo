package response

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int64  `json:"expireTime"`
	User        string `json:"user"`
}
