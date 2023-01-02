package data

type GrafanaUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	OrgId    int    `json:"OrgId"`
}
