package handlers

import "github.com/cookiengineer/golocron/config"
import "fmt"
import "log"
import "net/http"

func MethodNotAllowed(config *config.Config, request *http.Request, response http.ResponseWriter) {

	log.Println(fmt.Sprintf("> %s %s: %d", request.Method, request.URL.Path, http.StatusMethodNotAllowed))

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.WriteHeader(http.StatusMethodNotAllowed)
	response.Write([]byte("That technique is forbidden... and you were foolish to attempt it."))

}
