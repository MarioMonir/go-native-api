package json_utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ToJson(w http.ResponseWriter, T any) error {
	return json.NewEncoder(w).Encode(T)
}

func FromJSON(r io.Reader, T any) error {
	return json.NewDecoder(r).Decode(T)
}

func ReciveJsonWithHandleHttpError(w http.ResponseWriter, r io.Reader, T any) {
	err := FromJSON(r, T)
	if err != nil {
		http.Error(w, "Unable to Decode json", http.StatusBadRequest)
	}
}

func SendJsonWithHandleHttpError(w http.ResponseWriter, T any) {
	err := ToJson(w, T)
	if err != nil {
		http.Error(w, "Unable to Encode json", http.StatusInternalServerError)
		return
	}
}
