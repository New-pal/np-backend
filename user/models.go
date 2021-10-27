package user

import (
	"fmt"
	"github.com/New-pal/np-backend/core"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;index;unique" valid:"email,required" json:"email"`
	Password []byte `json:"-" swaggerignore:"true"`
	Name     string `json:"name"`
}

func (u *User) GetId() string {
	return fmt.Sprint(u.ID)
}
func (u *User) SetPassword(pwd string) error {
	password, err := core.HashPassword([]byte(pwd))
	if err != nil {
		return err
	}
	u.Password = password
	return nil
}
func (u *User) GetPassword() []byte {
	return u.Password
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetName() string {
	return u.Name
}
