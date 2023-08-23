package models

type User struct {
	Login        string
	PasswordHash string
	Name         string
	Age          int64
}
