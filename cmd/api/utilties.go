package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	// wrapper := make(map[string]interface{})
	// wrapper[wrap] = data

	js, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

type jsonError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (app *application) errorJson(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := jsonError{
		Status:  "error",
		Message: err.Error(),
	}

	app.writeJSON(w, statusCode, payload)
}
