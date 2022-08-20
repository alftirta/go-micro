package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	// Limit the size of incoming request body.
	maxBytes := 1048576 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Decode the request body into 'data'.
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	// Check if there is only 1 JSON value by trying to decode the next JSON-encoded value into an arbitrary variable, and it should return 'EOF' error.
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("the request body must have only a single JSON value")
	}

	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	// Encode 'data' into a slice of bytes ([]byte)
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// If 'headers' is specified, set the response header accordingly.
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// Set the "Content-Type" header to "application/json".
	w.Header().Set("Content-Type", "application/json")

	// Set the response status code according to the 'status' argument.
	w.WriteHeader(status)

	// Write the JSON response.
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	// Set the default status code.
	statusCode := http.StatusBadRequest

	// If 'status' is specified, set the status code accordingly.
	if len(status) > 0 {
		statusCode = status[0]
	}

	// Write the JSON response.
	return app.writeJSON(w, statusCode, jsonResponse{Error: true, Message: err.Error()})
}
