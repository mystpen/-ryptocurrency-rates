package delivery

import (
	"encoding/json"
	"log"
	"net/http"
)

type envelope map[string]any

func (h *Handler) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (h *Handler) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := h.writeJSON(w, status, env, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
}