package user

import (
	"errors"
	"gorm.io/gorm"
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
