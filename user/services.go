package user

import (
	"github.com/New-pal/np-backend/core"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) (*User, error) {
	usr, ok := c.Get("user")
	if !ok {
		return nil, core.ErrUnauthorized
	}
	u, ok := usr.(*User)
	if !ok {
		return nil, core.ErrUnauthorized
	}
	return u, nil
}

type UserService struct {
	con *gorm.DB
}

func (us *UserService) CreateUser(email string, firstName string, lastName string, gender bool, age uint,
	password string) (interface{ userInterface }, error) {
	user := User{Email: email, FirstName: firstName, LastName: lastName, Gender: gender, Age: age}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	res := us.con.Create(&user)

	return &user, res.Error
}

func (us *UserService) UpdateUser(user *User) error {
	return us.con.Save(user).Error
}

func NewUserService(con *gorm.DB) *UserService {
	return &UserService{con: con}
}
