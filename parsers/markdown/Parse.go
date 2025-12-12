package markdown

import net_url "net/url"
import "strings"

func Parse(base_url *net_url.URL, path string, bytes []byte) *Document {

	if base_url != nil && strings.HasPrefix(path, "/") {

		path_url, err := net_url.Parse(path)

		if err == nil {

			var document Document

			document.URL = base_url.ResolveReference(path_url)
			document.Meta.Tags = make([]string, 0)
			document.Meta.Author = "Golocron"
			document.Body = make([]*Element, 0)

			// TODO: Make this /path/to/teaser.jpg
			document.Meta.Image = "/teaser.jpg"

			markdown_code := strings.TrimSpace(string(bytes))

			if markdown_code != "" {
				document.ParseMeta(markdown_code)
				document.ParseBody(markdown_code)
			}

			return &document

		}

	}

	return nil

}
