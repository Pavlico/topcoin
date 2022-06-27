package response

import (
	"net/http"
)

type ResponseWriter interface {
	WriteResponse(w *http.ResponseWriter, status int, result []byte)
}

func writeHeaders(w *http.ResponseWriter, status int) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
}

func writeBody(w *http.ResponseWriter, result []byte) {
	_, err := (*w).Write(result)
	if err != nil {
		writeHeaders(w, http.StatusInternalServerError)
	}
}

func WriteResponse(w *http.ResponseWriter, status int, result []byte) {
	writeHeaders(w, status)
	writeBody(w, result)
}
