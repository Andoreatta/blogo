package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Returns a JSON response for a request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

// Returns an error in JSON format
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		//Bruh
		Error: err.Error(),
	})
}
