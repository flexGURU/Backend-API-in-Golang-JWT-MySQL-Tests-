package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {

	if r.Body == nil {
		return errors.New("empty request body")
	}
	
	return json.NewDecoder(r.Body).Decode(payload)
	
}

func WriteJSON(w http.ResponseWriter, payload any) error {
	return json.NewEncoder(w).Encode(payload)
}

