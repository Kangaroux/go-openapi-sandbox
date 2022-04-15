package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data interface{}, statusCode ...int) {
	serialized, err := json.Marshal(data)

	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	if len(statusCode) > 0 {
		w.WriteHeader(statusCode[0])
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(serialized)
}
