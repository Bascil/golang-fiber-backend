package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id uint `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`//array of bytes / do not show password
	RoleId uint `json:"role_id"`
	Role Role `json:"role" gorm:"foreignKey:RoleId"`
}

func (user *User) SetPassword(password string){
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error{
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}