package handlers

import "github.com/cookiengineer/golocron/config"
import "net/http"

func SeeOther(config *config.Config, request *http.Request, response http.ResponseWriter, location string) {

	response.Header().Set("Location", location)
	response.WriteHeader(http.StatusSeeOther)
	response.Write([]byte("Your answer lies elsewhere... follow the path I have chosen for you."))

}

