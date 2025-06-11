package models

type User struct {
	Username     string `json:"Логин"`
	Password     string `json:"Пароль"`
	DeletionMark bool   `json:"Пометка удаления"`
}

func (User) TableName() string {
	return "users"
}
