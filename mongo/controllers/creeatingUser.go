package controller

import (
	"encoding/json"
	"net/http"
	"os/user"
)

func CreatingUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user user.User
	_ = json.NewDecoder(r.Body).Decode((&user))
	inserData(user)
	json.NewEncoder(w).Encode(user)
}
