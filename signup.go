package order

import (
	"encoding/json"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Unexpected message format", http.StatusBadRequest)
		return
	}

	// TODO: Finish implementing this.
}
