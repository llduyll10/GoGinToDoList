package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func CheckPassword(hashPassword string, plainPassword []byte) (bool, error) {
	hashPw := []byte(hashPassword)
	if err := bcrypt.CompareHashAndPassword(hashPw, plainPassword); err != nil {
		return false, err
	}
	return true, nil
}
