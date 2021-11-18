package user

import (
	"fmt"
	"github.com/New-pal/np-backend/core"
	"gorm.io/gorm"
	"time"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Email     string     `json:"email" gorm:"size:255;not null;index;unique" valid:"email,required"`
	Password  []byte     `json:"-" swaggerignore:"true"`
	FirstName string     `json:"first_name" gorm:"size:255; not null;" valid:"required"`
	LastName  string     `json:"last_name" gorm:"size:255"`
	Gender    bool       `json:"gender" gorm:"type:bool;not null"` // false - female true - male
	Birthday  *time.Time `json:"birthday" gorm:"type:time;default=null"`
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
