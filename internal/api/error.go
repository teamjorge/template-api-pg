package api

import (
	"fmt"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, err error, code int, message string) {
	res := fmt.Sprintf(`{"error": "%s"}`, message)
	if code == 500 {
		log.Printf("%s - %v\n", message, err)
	}
	w.WriteHeader(code)
	w.Write([]byte(res))
}
