package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime"
	"net/http"

	schema "github.com/Kangaroux/go-map-schema"
)

// ParseRequestJSON unmarshals the JSON body from the request into `out`. If the
// request is not valid JSON, it sends a 400 error and returns false.
func ParseRequestJSON(w http.ResponseWriter, req *http.Request, out interface{}) bool {
	ct, _, err := mime.ParseMediaType(req.Header.Get("content-type"))

	// Check the content type
	if err != nil || ct != "application/json" {
		WriteJSON(w, NewErrorResponse("content type must be application/json"), 400)
		return false
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Print(err)
		WriteJSON(w, NewInternalErrorResponse(), 500)
		return false
	}

	m := make(map[string]interface{})

	// Load the request into a map so we can check the schema
	if err = json.Unmarshal(body, &m); err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			log.Print(err)
			WriteJSON(w, NewErrorResponse("request body must be a json object"), 400)
			return false
		}

		log.Print(err)
		WriteJSON(w, NewErrorResponse("failed to parse json: "+err.Error()), 400)
		return false
	}

	results, err := schema.CompareMapToStruct(out, m, &schema.CompareOpts{TypeNameFunc: schema.SimpleTypeName})

	if err != nil {
		log.Print(err)
		WriteJSON(w, NewInternalErrorResponse(), 500)
		return false
	}

	fieldErrors := results.Errors()

	// Return any field errors
	if fieldErrors != nil {
		WriteJSON(w, NewFieldErrorResponse(fieldErrors, "type mismatch on one or more fields"), 400)
		return false
	}

	// Should be able to safely load the data into the destination now
	if err = json.Unmarshal(body, &out); err != nil {
		log.Print(err)
		WriteJSON(w, NewErrorResponse("failed to parse json: "+err.Error()), 400)
		return false
	}

	return true
}

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
