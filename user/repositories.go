package user

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetUserById(id string) (*User, error) {
	if len(id) == 0 {
		return nil, errors.New("empty id")
	}
	var u User
	err := repo.db.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &u, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (interface{ userInterface }, error) {
	if len(email) == 0 {
		return nil, errors.New("empty email")
	}
	var u User
	err := repo.db.Where(&User{Email: email}).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &u, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserSettingsRepository struct {
	db *gorm.DB
}

func (repo *UserSettingsRepository) GetSettingsById(id string) (*UserSettings, error) {
	if len(id) == 0 {
		return nil, errors.New("empty id")
	}
	var settings UserSettings
	err := repo.db.First(&settings, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("settings not found")
	}
	return &settings, nil
}

func (repo *UserSettingsRepository) GetSettingsByUserId(userId string) (*UserSettings, error) {
	if len(userId) == 0 {
		return nil, errors.New("empty id")
	}
	id, convErr := strconv.ParseUint(userId, 10, 32)
	if convErr != nil {
		return nil, errors.New("incorrect user ID string")
	}
	var settings UserSettings
	err := repo.db.Where(&UserSettings{UserId: uint(id)}).First(&settings).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("settings not found")
	}
	return &settings, nil
}

func NewUserSettingsRepository(db *gorm.DB) *UserSettingsRepository {
	return &UserSettingsRepository{db: db}
}
