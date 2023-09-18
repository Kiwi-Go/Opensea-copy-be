package models

type Account struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	NAME     string `json:"name"`
	PASSWORD string `json:"password"`
}
