package auth

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`

	Claims map[string]string `json:"claims,omitempty"`
}
