package templates

import net_url "net/url"

func RenderURL(base_url *net_url.URL, ref string) string {

	var result string

	if ref == "" {

		result = base_url.String()

	} else {

		ref_url, err := net_url.Parse(ref)

		if err == nil {
			result = base_url.ResolveReference(ref_url).String()
		}

	}

	return result

}
