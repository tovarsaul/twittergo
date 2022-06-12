package routers

import (
	"encoding/json"
	"net/http"
	"twittergo/domain/profile"
)

func SearchProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must sent id param", http.StatusBadRequest)
		return
	}
	user, err := profile.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Error searching user "+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
