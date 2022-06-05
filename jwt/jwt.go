package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"twittergo/models"
)

func GenerateJWT(user models.User) (string, error) {
	key := []byte("DevelopmentMasters")
	payload := jwt.MapClaims{
		"mail":      user.Mail,
		"name":      user.Name,
		"lastname":  user.LastName,
		"birthday":  user.Birthday,
		"biography": user.Biography,
		"location":  user.Location,
		"website":   user.Website,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return signedToken, err
	}
	return signedToken, err
}
