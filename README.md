
# golocron

Pluggable Blogging and Wiki System for Go Backends.

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
Integration of filepaths and prefixes can be done on a per-prefix basis. All Editor relevant web assets
are integrated with an `embed.FS` instance so that you don't have to worry about mapping those paths.

Default mapping paths:

- `/golocron/index.html` serves the Editor UI
- `/golocron/api/index` serves the document index for published markdown documents
- `/golocron/api/open` opens markdown documents
- `/golocron/api/save` saves markdown documents

### Integration of Golocron Editor

TODO: Example

### Live Preview Integration

It is assumed that documents are rendered with their file extension, meaning that a document that is
located at `/weblog/1337/example.md` is rendered by the local webserver at the HTTP path `/weblog/1337/example.html`.

This will lead to a nice preview feature, where the right side of the Editor UI will display the
rendered markdown articles as their HTML previews.

### Template Rendering

If you want to use custom templates, it is recommended to use the [parsers/markdown.Parse](./parsers/markdown/Parse.go)
method, and to use the `document.Render()` method to generate an indented HTML string.


### License

[AGPL-3.0](./LICENSE.txt)

