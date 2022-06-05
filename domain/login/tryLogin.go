package login

import (
	"golang.org/x/crypto/bcrypt"
	"twittergo/domain/register"
	"twittergo/models"
)

func TryLogin(mail string, password string) (models.User, bool) {
	user, find, _ := register.CheckIfUserExists(mail)
	if !find {
		return user, find
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, find
}
