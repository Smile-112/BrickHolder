package models

type User struct {
	Username     string `json:"Логин" gorm:"column:username; primary_key"`
	Password     string `json:"Пароль" gorm:"column:password"`
	DeletionMark bool   `json:"Пометка удаления" gorm:"column:deletionmark"`
}

func (User) TableName() string {
	return "users"
}
