package api

import "github.com/cookiengineer/golocron/config"
import "github.com/cookiengineer/golocron/handlers"
import "github.com/cookiengineer/golocron/parsers/markdown"
import "github.com/cookiengineer/golocron/templates"
import "net/http"
import "os"
import "strings"

func Render(config *config.Config, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		file := request.URL.Query().Get("file")

		if strings.HasPrefix(file, "/") && strings.HasSuffix(file, ".md") {

			base_url  := config.BaseURL()
			file_path := config.ResolvePath(file)

			bytes, err1 := os.ReadFile(file_path)

			if err1 == nil {

				document := markdown.Parse(base_url, file, bytes)

				if document != nil {
					handlers.ServeTemplate(config, request, response, templates.DocumentTemplate, document)
				} else {
					handlers.InternalServerError(config, request, response)
				}

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
