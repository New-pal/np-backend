package core

import "golang.org/x/crypto/bcrypt"

func CheckPassword(hash []byte, pwd string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(pwd))
}

func HashPassword(pwd []byte) ([]byte, error) {
	if len(pwd) == 0 {
		return []byte{}, errEmptyPassword
	}
	return bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
}
