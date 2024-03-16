package json_utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func SendJson(w http.ResponseWriter, i interface{}) {
	err := json.NewEncoder(w).Encode(i)
	if err != nil {
		http.Error(w, "Unable to Encode json", http.StatusInternalServerError)
	}
}

func RecieveJson(w http.ResponseWriter, r io.Reader, i interface{}) {
	e := json.NewDecoder(r)
	err := e.Decode(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
