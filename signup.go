package signup

import (
	"encoding/json"
	"log"
	"net/http"
)

func EmailSignup(w http.ResponseWriter, r *http.Request) {
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
		Email    string `json:"email"`
		Honeypot string `json:"ninja"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Unexpected message format", http.StatusBadRequest)
		return
	}

	// Real users should never submit anything in the honeypot field, as it's only
	// visible to bots.
	if payload.Honeypot != "" {
		log.Printf("bot signup detected: %s", payload.Email)
		// Return a fake success message
		return
	}

	// Don't bother validating because the upstream server will validate for us.
	err := addSubscriber(payload.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
