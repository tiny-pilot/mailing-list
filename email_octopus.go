package signup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func addSubscriber(emailAddress string) error {
	apiKey := mustGetenv("EMAIL_OCTOPUS_API_KEY")
	listId := mustGetenv("EMAIL_OCTOPUS_LIST_ID")

	url := fmt.Sprintf("https://emailoctopus.com/api/1.5/lists/%s/contacts", listId)
	postBody, _ := json.Marshal(map[string]string{
		"api_key":       apiKey,
		"email_address": emailAddress,
	})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var respBody struct {
			Error struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"error"`
		}
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			return err
		}
		log.Printf("Failed to add subscriber: %s - %s", respBody.Error.Code, respBody.Error.Message)
		return fmt.Errorf("Failed to add subscriber: %s", respBody.Error.Message)
	}

	var respBody struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return err
	}
	return nil
}

func mustGetenv(key string) string {
	e := os.Getenv(key)
	if e == "" {
		log.Fatalf("%s environment variable must be set", key)
	}
	return e
}
