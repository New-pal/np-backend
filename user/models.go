package user

import (
	"fmt"
	"github.com/New-pal/np-backend/core"
	"gorm.io/gorm"
	"strconv"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"size:255;not null;index;unique" valid:"email,required"`
	Password  []byte `json:"-" swaggerignore:"true"`
	FirstName string `json:"first_name" gorm:"size:255; not null;" valid:"required"`
	LastName  string `json:"last_name" gorm:"size:255"`
	Gender    bool   `json:"gender" gorm:"type:bool;not null"` // false - female true - male
	Age       uint   `json:"age" gorm:"type:uint;not null"`
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

type UserSettings struct {
	gorm.Model
	UserId             uint
	User               User   `json:"user" gorm:"constraint: OnDelete:CASCADE; default:null"`
	Language           string `json:"language" gorm:"type:varchar(2)"`
	LastNameVisibility bool   `json:"last_name_visibility" gorm:"type:bool; default:true"`
	AgeVisibility      bool   `json:"age_visibility" gorm:"type:bool; default:true"`
	TimeFormat         bool   `json:"time_format" gorm:"type:bool; default:true"`
	DistanceUnits      bool   `json:"distance_units" gorm:"type:bool; default:true"`
}

func (us *UserSettings) GetUserId() string {
	return strconv.Itoa(int(us.UserId))
}
func (us *UserSettings) GetLang() string {
	return us.Language
}
func (us *UserSettings) GetLastnameVisibility() bool {
	return us.LastNameVisibility
}
func (us *UserSettings) GetAgeVisibility() bool {
	return us.AgeVisibility
}
func (us *UserSettings) GetTimeFormat() bool {
	return us.TimeFormat
}
func (us *UserSettings) GetDistanceUnits() bool {
	return us.DistanceUnits
}
