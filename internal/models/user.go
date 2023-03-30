package models

type User struct {
	Id       int
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}
