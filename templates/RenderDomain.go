package templates

import net_url "net/url"
import "strings"

func RenderDomain(url *net_url.URL) string {

	var result string

	if strings.Contains(url.Host, ":") {

		return url.Host[0:strings.Index(url.Host, ":")]

	} else {
		return url.Host
	}

	return result

}
