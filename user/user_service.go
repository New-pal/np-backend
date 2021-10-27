package user

import (
	"gorm.io/gorm"
)

type UserService struct {
	con *gorm.DB
}

func (us *UserService) CreateUser(email string, name string, password string) (interface{ userInterface }, error) {
	user := User{Email: email, Name: name}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	res := us.con.Create(&user)

	return &user, res.Error
}

func NewUserService(con *gorm.DB) *UserService {
	return &UserService{con: con}
}
