package data

type HiveosAuth struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	TwofaCode string `json:"twofa_code"`
	Remember  bool   `json:"remember"`
}

type HiveosToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
