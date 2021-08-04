package order

import (
	"encoding/json"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://tinypilotkvm.com")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Only OPTIONS and POST are supported", http.StatusMethodNotAllowed)
	}

	var payload struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Unexpected message format", http.StatusBadRequest)
		return
	}

	// Don't bother validating because the upstream server will validate for us.
	// subscriberEmail := payload.Email

	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
