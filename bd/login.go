package bd

import (
	"github.com/KevinDanae/twitterClone/models"
	"golang.org/x/crypto/bcrypt"
)

// Login is a function that logs a user in to the database
func Login(email string, password string) (models.User, bool) {
	us, exist, _ := CheckUserExist(email)
	if !exist {
		return us, false
	}

	passwordB := []byte(password)
	passwordBD := []byte(us.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordB)
	if err != nil {
		return us, false
	}

	return us, true
}
