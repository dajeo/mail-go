package models

type JWT struct {
	Token    string `json:"token"`
	Username string `json:"user_nicename"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRes struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
