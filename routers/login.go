package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"twittergo/domain/login"
	"twittergo/jwt"
	"twittergo/models"
)

func Login(writer http.ResponseWriter, response *http.Request) {
	writer.Header().Add("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "User/password not valid "+err.Error(), 400)
		return
	}
	if len(user.Mail) == 0 {
		http.Error(writer, "User mail is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(writer, "User password is to short", 400)
		return
	}
	document, exist := login.TryLogin(user.Mail, user.Password)
	if !exist {
		http.Error(writer, "User/password not valid", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(writer, "Error trying to generate token "+err.Error(), 400)
		return
	}
	resp := models.ResponseLogin{
		Token: jwtKey,
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resp)

	//Guardar la cockie del usuario
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
