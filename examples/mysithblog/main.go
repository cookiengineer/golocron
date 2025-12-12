package main

import "github.com/cookiengineer/golocron"
import golocron_config "github.com/cookiengineer/golocron/config"
import golocron_handlers "github.com/cookiengineer/golocron/handlers"
import "github.com/cookiengineer/golocron/parsers/markdown"
import "net/http"
import "os"
import "strings"
import "fmt"

func main() {

	cwd, _ := os.Getwd()

	// Base URL is http://localhost:3000
	// Live-Preview enabled
	config := golocron_config.NewConfig("http://localhost:3000", false, map[string]string{
		"/weblog": cwd + "/public/weblog",
		"/wiki":   cwd + "/public/wiki",
	})

	mux := http.NewServeMux()

	// Your custom existing Routes
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		if request.Method == http.MethodGet && strings.HasPrefix(request.URL.Path, "/") {

			if request.URL.Path == "/" {
				golocron_handlers.SeeOther(config, request, response, "/index.html")
			} else if request.URL.Path == "/index.html" {

				response.Header().Set("Content-Type", "text/html; charset=utf-8")

				html := `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Darth Sidious' Blog</title>
	</head>
	<body>
		<h1>Welcome to my Blog</h1>
		<p>
			At last... the <a href="/golocron/index.html">Golocron</a>. The power of the Sith, all in your hands.
			And here I thought I was clever - turns out, I'm just getting started.
			Ha ha ha.
		</p>
	</body>
</html>
`

				response.Write([]byte(html))

			} else if strings.HasPrefix(request.URL.Path, "/weblog/") || strings.HasPrefix(request.URL.Path, "/wiki/") {

				if strings.HasSuffix(request.URL.Path, ".md") {

					// The Sith never share their source of power!
					golocron_handlers.NotFound(config, request, response)

				} else if strings.HasSuffix(request.URL.Path, ".html") {

					// Live Preview
					http_path := strings.TrimSuffix(request.URL.Path, ".html") + ".md"
					file_path := config.ResolvePath(http_path)

					buffer, err := os.ReadFile(file_path)

					if err == nil {

						document := markdown.Parse(config.BaseURL(), http_path, buffer)

						html := fmt.Sprintf(`<!DOCTYPE html>
<html>
	<head>
		<title>%s</title>
	</head>
	<body>
		<article>
%s
		</article>
	</body>
</html>
`, document.Meta.Title, document.Render("\t\t\t"))

						golocron_handlers.ServeFile(config, request, response, []byte(html))

					} else {
						golocron_handlers.NotFound(config, request, response)
					}

				} else {
					golocron_handlers.NotFound(config, request, response)
				}

			} else {
				golocron_handlers.NotFound(config, request, response)
			}

		} else {
			golocron_handlers.MethodNotAllowed(config, request, response)
		}

	})

	// Dispatch golocron to /golocron/ and /golocron/api
	golocron.Activate(config, mux)

	http.ListenAndServe(":3000", mux)

}
