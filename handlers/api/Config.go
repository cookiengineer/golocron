package api

import "github.com/cookiengineer/golocron/config"
import "github.com/cookiengineer/golocron/handlers"
import "github.com/cookiengineer/golocron/parsers/markdown"
import "encoding/json"
import "net/http"
import io_fs "io/fs"
import "os"
import "strconv"
import "strings"

type Handshake struct {
	BaseURL     string            `json:"base_url"`
	LivePreview bool              `json:"live_preview"`
	Documents   map[string]string `json:"documents"`
}

func Config(config *config.Config, request *http.Request, response http.ResponseWriter) {

	if request.Method == http.MethodGet {

		base_url  := config.BaseURL()
		handshake := Handshake{
			BaseURL:     base_url.String(),
			LivePreview: config.LivePreview(),
			Documents:   make(map[string]string),
		}

		document_paths := config.GetPaths()

		for _, path := range document_paths {

			fs            := os.DirFS(config.GetRoot(path))
			entries, err0 := io_fs.ReadDir(fs, ".")

			if err0 == nil {

				for _, entry := range entries {

					filename := entry.Name()

					if strings.HasSuffix(filename, ".md") && entry.IsDir() == false {

						bytes, err1 := io_fs.ReadFile(fs, filename)

						if err1 == nil {

							document := markdown.Parse(base_url, path + "/" + filename, bytes)

							if document != nil && document.IsValid() {
								handshake.Documents[document.URL.Path] = document.Meta.Date.Format("2006-01-02")
							} else {
								handshake.Documents[document.URL.Path] = "0000-00-00"
							}

						}

					}

				}

			}

		}

		buffer, err := json.MarshalIndent(handshake, "", "\t")

		if err == nil {

			response.Header().Set("Content-Type", "application/json")
			response.Header().Set("Content-Encoding", "identity")
			response.Header().Set("Content-Length", strconv.Itoa(len(buffer)))
			response.Write(buffer)

		} else {
			handlers.InternalServerError(config, request, response)
		}

	} else {
		handlers.MethodNotAllowed(config, request, response)
	}

}
