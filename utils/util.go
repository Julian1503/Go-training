package Util

import (
	"encoding/json"
	"log"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "Application/json")
	log.Fatal(json.NewEncoder(w).Encode(data))
}
