package model

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"` // store hashed password in production!
}
