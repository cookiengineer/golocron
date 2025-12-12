package handlers

import "github.com/cookiengineer/golocron/config"
import "fmt"
import "log"
import "net/http"

func InternalServerError(config *config.Config, request *http.Request, response http.ResponseWriter) {

	log.Println(fmt.Sprintf("> %s %s: %d", request.Method, request.URL.Path, http.StatusInternalServerError))

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte("The system trembles under its own weakness. Let the failure consume it."))

}
