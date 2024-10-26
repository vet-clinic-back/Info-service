package models

type Owner struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"` // password hash
}
