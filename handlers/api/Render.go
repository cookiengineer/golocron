package api

import "github.com/cookiengineer/golocron/config"
import "github.com/cookiengineer/golocron/handlers"
import "github.com/cookiengineer/golocron/parsers/markdown"
import "github.com/cookiengineer/golocron/templates"
import "fmt"
import "log"
import "net/http"
import "net/url"
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

				path_url, err2 := url.Parse(file)

				if err2 == nil {

					document_url := base_url.ResolveReference(path_url)
					document, err3 := markdown.Parse(document_url.String(), bytes)

					if err3 == nil {
						handlers.ServeTemplate(config, request, response, templates.DocumentTemplate, document)
					} else {
						log.Println(fmt.Sprintf("> %s %s: %d", request.Method, request.URL.Path, err3.Error()))
						handlers.InternalServerError(config, request, response)
					}

				} else {
					log.Println(fmt.Sprintf("> %s %s: %d", request.Method, request.URL.Path, err2.Error()))
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
