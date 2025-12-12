package handlers

import "github.com/cookiengineer/golocron/config"
import io_fs "io/fs"
import "net/http"
import "strings"

func ServeFilesystem(config *config.Config, request *http.Request, response http.ResponseWriter, fs io_fs.FS) {

	if request.Method == http.MethodGet {

		bytes, err := io_fs.ReadFile(fs, strings.TrimPrefix(request.URL.Path, "/"))

		if err == nil {
			ServeFile(config, request, response, bytes)
		} else {
			NotFound(config, request, response)
		}

	} else {
		MethodNotAllowed(config, request, response)
	}

}
