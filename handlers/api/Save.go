package api

import "github.com/cookiengineer/golocron/handlers"
import "github.com/cookiengineer/golocron/config"
import "fmt"
import "log"
import "io"
import "net/http"
import "os"
import "strings"

func Save(config *config.Config, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodPost {

		save_as      := request.Header.Get("X-Save-As")
		content_type := request.Header.Get("Content-Type")

		if content_type == "text/markdown" && strings.HasPrefix(save_as, "/") && strings.HasSuffix(save_as, ".md") {

			// base_url  := config.BaseURL()
			file_path := config.ResolvePath(save_as)

			bytes, err0 := io.ReadAll(request.Body)

			if err0 == nil {

				err1 := os.WriteFile(file_path, bytes, 0644)

				if err1 == nil {

					log.Println(fmt.Sprintf("> %s %s: %s %d", request.Method, request.URL.Path, save_as, http.StatusOK))

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusOK)
					response.Write([]byte(""))

				} else {

					log.Println(fmt.Sprintf("> %s %s: %s %d", request.Method, request.URL.Path, save_as, http.StatusInternalServerError))

					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusInternalServerError)
					response.Write([]byte(""))

				}

			} else {

				log.Println(fmt.Sprintf("> %s %s: %s %d", request.Method, request.URL.Path, save_as, http.StatusInternalServerError))

				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusInternalServerError)
				response.Write([]byte(""))

			}

		} else {
			handlers.UnsupportedMediaType(config, request, response)
		}

	} else {
		handlers.MethodNotAllowed(config, request, response)
	}

}
