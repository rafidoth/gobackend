package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Envelope map[string]interface{}

func WriteJson(w http.ResponseWriter, status int, data Envelope) error {
	js, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func ReadIntParam(r *http.Request, paramName string) (int64, error) {
	fetchedParam := chi.URLParam(r, paramName)
	if fetchedParam == "" {
		return 0, errors.New("invalid param")
	}

	intParam, err := strconv.ParseInt(fetchedParam, 10, 64)
	if err != nil {
		return 0, errors.New("param is not a valid integer")
	}

	return intParam, nil
}
