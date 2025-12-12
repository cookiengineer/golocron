package handlers

import "github.com/cookiengineer/golocron/config"
import "fmt"
import "log"
import "net/http"

func NotFound(config *config.Config, request *http.Request, response http.ResponseWriter) {

	log.Println(fmt.Sprintf("> %s %s: %d", request.Method, request.URL.Path, http.StatusNotFound))

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.WriteHeader(http.StatusNotFound)
	response.Write([]byte("What you seek has vanished into the void. Only shadows remain."))

}
