package routers

import (
	"encoding/json"
	"net/http"
	"twittergo/domain/register"
	"twittergo/models"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Error deserialized user model received "+err.Error(), 400)
		return
	}
	if len(user.Mail) == 0 {
		http.Error(writer, "Mail is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(writer, "Password is to short", 400)
		return
	}
	_, find, _ := register.CheckIfUserExists(user.Mail)
	if find {
		http.Error(writer, "Mail exists", 400)
		return
	}
	_, status, err := register.InsertRegister(user)
	if err != nil {
		http.Error(writer, "Error occurs while registry user "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(writer, "User cannot be register"+err.Error(), 400)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
