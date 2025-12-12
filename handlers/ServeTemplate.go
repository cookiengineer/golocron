package handlers

import "github.com/cookiengineer/golocron/config"
import "github.com/cookiengineer/golocron/parsers/markdown"
import "bytes"
import "compress/gzip"
import "net/http"
import "strconv"
import "strings"
import "text/template"

func ServeTemplate(config *config.Config, request *http.Request, response http.ResponseWriter, template *template.Template, document *markdown.Document) {

	if request.Method == http.MethodGet {

		if template != nil && document.IsValid() == true {

			var buffer bytes.Buffer

			err1 := template.Execute(&buffer, document)

			if err1 == nil {

				policies := []string{

					// Assets
					"default-src 'self'",
					"script-src 'self'",
					"script-src-elem 'self'",
					"font-src 'self'",
					"img-src 'self'",
					"style-src 'self'",

					// Disallow Trackers
					"prefetch-src 'none'",
					"worker-src 'self'",

					// Allow XHR/Fetch to other websites
					"media-src 'self'",
					"frame-src * 'self'",
					"connect-src * 'self'",
				}

				for _, policy := range policies {
					response.Header().Set("Content-Security-Policy", policy)
				}

				accept_encoding := request.Header.Get("Accept-Encoding")

				if strings.Contains(accept_encoding, "gzip") {

					var encoded_buffer bytes.Buffer

					writer := gzip.NewWriter(&encoded_buffer)

					_, err1 := writer.Write(buffer.Bytes());

					if err1 == nil {

						err2 := writer.Close()

						if err2 == nil {

							encoded_bytes := encoded_buffer.Bytes()

							response.Header().Set("Content-Type", "text/html; charset=utf-8")
							response.Header().Set("Content-Encoding", "gzip")
							response.Header().Set("Content-Length", strconv.Itoa(len(encoded_bytes)))
							response.Write(encoded_bytes)

						} else {
							InternalServerError(config, request, response)
						}

					} else {
						InternalServerError(config, request, response)
					}

				} else {

					raw_bytes := buffer.Bytes()

					response.Header().Set("Content-Type", "text/html; charset=utf-8")
					response.Header().Set("Content-Encoding", "identity")
					response.Header().Set("Content-Length", strconv.Itoa(len(raw_bytes)))
					response.Write(raw_bytes)

				}

			} else {
				InternalServerError(config, request, response)
			}

		} else {
			InternalServerError(config, request, response)
		}

	} else {
		MethodNotAllowed(config, request, response)
	}

}
