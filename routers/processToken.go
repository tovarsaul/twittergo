package routers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"twittergo/domain/register"
	"twittergo/models"
)

var (
	mail, idUser string
)

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("DevelopmentMasters")
	claim := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claim, false, "", errors.New("token format invalid")
	}
	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, find, _ := register.CheckIfUserExists(claim.Mail)
		if find {
			mail = claim.Mail
			idUser = claim.Id
		}
		return claim, find, idUser, err
	}
	if !tkn.Valid {
		return claim, false, "", errors.New("invalid token")
	}
	return claim, false, "", err
}
