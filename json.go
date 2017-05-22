package httpcore

import (
	"encoding/json"
	"errors"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, content interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(content)
}

func WriteDataJSON(w http.ResponseWriter, status int, data Identifier) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(struct {
		Data interface{} `json:"data"`
		Type string      `json:"type"`
		ID   string      `json:"id"`
	}{
		Data: data,
		ID:   data.GetID(),
		Type: data.GetType(),
	})
}

func WriteErrorsJSON(w http.ResponseWriter, status int, errs ...error) {
	w.WriteHeader(status)
	if len(errs) == 0 || (len(errs) == 1 && errs[0] == nil) {
		errs = append(errs, errors.New(http.StatusText(status)))
	}
	err := json.NewEncoder(w).Encode(struct {
		Errs []error `json:"errors"`
	}{
		Errs: errs,
	})
	if err != nil {
		panic(err)
	}
}
