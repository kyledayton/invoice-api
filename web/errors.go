package web

import (
	"encoding/json"
	"net/http"
)

func errJson(err error) []byte {
	data, _ := json.Marshal(map[string]string{
		"error": err.Error(),
	})

	return data
}

func renderError(res http.ResponseWriter, err error, status int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(errJson(err))
}
