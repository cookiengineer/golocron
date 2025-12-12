package api

import "github.com/cookiengineer/golocron/config"
import "github.com/cookiengineer/golocron/handlers"
import "net/http"
import "os"
import "strconv"
import "strings"

func Open(config *config.Config, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		file := request.URL.Query().Get("file")

		if strings.HasPrefix(file, "/") && strings.HasSuffix(file, ".md") {

			// base_url  := config.BaseURL()
			file_path := config.ResolvePath(file)

			bytes, err1 := os.ReadFile(file_path)

			if err1 == nil {

				response.Header().Set("Content-Type", "text/markdown; charset=utf-8")
				response.Header().Set("Content-Encoding", "identity")
				response.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
				response.Write(bytes)

			} else {
				handlers.NotFound(config, request, response)
			}

		} else {
			handlers.NotFound(config, request, response)
		}

	} else {
		handlers.MethodNotAllowed(config, request, response)
	}

}
