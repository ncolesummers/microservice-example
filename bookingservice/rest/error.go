package rest

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Msg string `json:"msg"`
}

func respondWithError(res http.ResponseWriter, msg string, code int) error {
	response := errorResponse{msg}
	_, err := json.Marshal(&response)
	if err != nil {
		return err
	}

	return nil
}
