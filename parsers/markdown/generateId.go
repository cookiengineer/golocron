package markdown

import "github.com/cookiengineer/golocron/parsers/utils"
import "strings"

func generateId(element *Element) string {

	texts := make([]string, 0)

	if element.Text != "" {

		chunks := strings.Split(strings.TrimSpace(utils.ToASCII(element.Text)), "-")

		for _, chunk := range chunks {

			tmp := strings.TrimSpace(chunk)

			if tmp != "" {
				texts = append(texts, tmp)
			}

		}

	} else if len(element.Children) > 0 {

		for _, child := range element.Children {

			if child.Type == "b" || child.Type == "code" || child.Type == "del" || child.Type == "em" || child.Type == "#text" {

				tmp := strings.TrimSpace(utils.ToASCIIName(child.Text))

				if strings.HasPrefix(tmp, "-") {
					tmp = tmp[1:]
				}

				if strings.HasSuffix(tmp, "-") {
					tmp = tmp[0 : len(tmp)-1]
				}

				chunks := strings.Split(tmp, "-")

				for _, chunk := range chunks {

					tmp := strings.TrimSpace(chunk)

					if tmp != "" {
						texts = append(texts, tmp)
					}

				}

			}

		}

	}

	filtered := make([]string, 0)

	if len(texts) > 0 {

		if utils.IsNumber(string(texts[0][0])) {
			texts = texts[1:]
		}

		for _, text := range texts {

			tmp := strings.ToLower(strings.TrimSpace(text))

			if tmp != "" {
				filtered = append(filtered, tmp)
			}

		}

		return strings.Join(filtered, "-")

	}

	return ""

}
