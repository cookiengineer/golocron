
# golocron

Plug and Play Blogging and Wiki System for Go Backends.

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/golocron/master/assets/golocron.jpg"/>
</p>

### Features

- Markdown syntax highlighting support
- Pluggable API to integrate into existing Go Backends
- Pluggable UI to integrate into existing Go Backends
- Stores documents on the Filesystem

### Usage

The Editor UI can be integrated into existing backends, by using the predefined [handlers](./handlers).
Integration of filesystem root folders are done on a per-path-prefix basis. All Editor relevant web
assets are integrated with an `embed.FS` instance so that you don't have to worry about mapping those
paths, they're served automatically at `/golocron/*`.

Default routes:

- `/golocron/index.html` serves the Editor UI
- `/golocron/api/config` serves the config and document index
- `/golocron/api/open` opens markdown document
- `/golocron/api/save` saves markdown document

Routes used when live-preview is disabled:

- `/golocron/api/render` renders markdown document

The integration is made quite simple. Take a look at the [examples/mysithblog](./examples/mysithblog)
to learn how to use the `golocron.Activate(*config.Config, *http.ServeMux)` API.

```go
// Minimal Example
func main() {

	cwd, _ := os.Getwd()

	// Base URL is http://localhost:3000
	// Live-Preview disabled
	config := golocron_config.NewConfig("http://localhost:3000", false, map[string]string{
		"/weblog": cwd + "/public/weblog",
		"/wiki":   cwd + "/public/wiki",
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		// TODO: Integrate your own Website routes
	})

	// Dispatch golocron to /golocron/ and /golocron/api
	golocron.Activate(config, mux)

	http.ListenAndServe(":3000", mux)

}
```

### Live Preview Integration

There's two essential modes of using Golocron: with Live-Preview enabled or Live-Preview disabled.

If the Live-Preview is `enabled`, it is assumed that the same paths are rendered by your existing
routes. So for example, `/weblog/foo/bar-qux.md` must be available as `/weblog/foo/bar-qux.html`.

If the Live-Preview is `disabled`, the Golocron Editor will use the `/golocron/api/render` route
to render the markdown document with default stylesheets.

### Template Rendering

If you want to use custom templates, it is recommended to use the [parsers/markdown.Parse](./parsers/markdown/Parse.go)
method, and to use the `document.Render()` method to generate an indented HTML string.

Golocron also offers an integration with this, take a look at the [templates](./templates)
folder which shows the default `text/template` instance that is used by the `/golocron/api/render`
route.


### License

[AGPL-3.0](./LICENSE.txt)

