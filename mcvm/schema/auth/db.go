package auth

type AuthUser struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	UUID     string `json:"uuid,omitempty"`
}
