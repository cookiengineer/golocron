package handlers

import "github.com/cookiengineer/golocron/config"
import "fmt"
import "log"
import "net/http"

func UnsupportedMediaType(config *config.Config, request *http.Request, response http.ResponseWriter) {

	log.Println(fmt.Sprintf("> %s %s: %d", request.Method, request.URL.Path, http.StatusUnsupportedMediaType))

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.WriteHeader(http.StatusUnsupportedMediaType)
	response.Write([]byte("Your data fails you... as you have failed the Sith. Correct your format, or be discarded."))

}
