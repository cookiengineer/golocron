package golocron

import "github.com/cookiengineer/golocron/handlers"
import "github.com/cookiengineer/golocron/handlers/api"
import "github.com/cookiengineer/golocron/config"
import io_fs "io/fs"
import "net/http"
import "strings"

func Activate(config *config.Config, mux *http.ServeMux) bool {

	fs, err := io_fs.Sub(Filesystem, "public")

	if err == nil {

		mux.HandleFunc("/golocron/", func(response http.ResponseWriter, request *http.Request) {

			if request.Method == http.MethodGet && strings.HasPrefix(request.URL.Path, "/") {

				if request.URL.Path == "/golocron/" {
					handlers.SeeOther(config, request, response, "/golocron/index.html")
				} else {
					handlers.ServeFilesystem(config, request, response, fs)
				}

			} else {
				handlers.MethodNotAllowed(config, request, response)
			}

		})

		mux.HandleFunc("/golocron/api/config", func(response http.ResponseWriter, request *http.Request) {

			if request.Method == http.MethodGet {
				api.Config(config, request, response)
			} else {
				handlers.MethodNotAllowed(config, request, response)
			}

		})

		mux.HandleFunc("/golocron/api/open", func(response http.ResponseWriter, request *http.Request) {

			if request.Method == http.MethodGet {
				api.Open(config, request, response)
			} else {
				handlers.MethodNotAllowed(config, request, response)
			}

		})

		mux.HandleFunc("/golocron/api/render", func(response http.ResponseWriter, request *http.Request) {

			if request.Method == http.MethodGet {
				api.Render(config, request, response)
			} else {
				handlers.MethodNotAllowed(config, request, response)
			}

		})

		mux.HandleFunc("/golocron/api/save", func(response http.ResponseWriter, request *http.Request) {

			if request.Method == http.MethodPost {

				save_as      := request.Header.Get("X-Save-As")
				content_type := request.Header.Get("Content-Type")

				if content_type == "text/markdown" && strings.HasPrefix(save_as, "/") && strings.HasSuffix(save_as, ".md") {
					api.Save(config, request, response)
				} else {
					handlers.MethodNotAllowed(config, request, response)
				}

			} else {
				handlers.MethodNotAllowed(config, request, response)
			}

		})

		return true

	}

	return false

}
