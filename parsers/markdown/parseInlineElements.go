package markdown

import "fmt"
import "strings"

func getTokenPosition(value string, remaining string) int {
	return len(strings.TrimSpace(value)) - len(remaining)
}

func parseInlineElements(value string, line int) ([]*Element, error) {

	result := make([]*Element, 0)
	remaining := strings.TrimSpace(value)

	for len(remaining) > 0 {

		is_emoji := strings.HasPrefix(remaining, ":")
		is_media := strings.HasPrefix(remaining, "![") && strings.Contains(remaining, "](") && strings.Contains(remaining, ")")
		is_a := strings.HasPrefix(remaining, "[") && strings.Contains(remaining, "](") && strings.Contains(remaining, ")")
		is_footnote := strings.HasPrefix(remaining, "[#") && strings.Contains(remaining, "]")
		is_abbr := strings.HasPrefix(remaining, "[") && strings.Contains(remaining, "]{") && strings.Contains(remaining, "}")
		is_b := strings.HasPrefix(remaining, "**")
		is_code := strings.HasPrefix(remaining, "`")
		is_em := strings.HasPrefix(remaining, "*")
		is_del := strings.HasPrefix(remaining, "~")

		if is_emoji {

			remaining = remaining[1:]

			index2 := strings.Index(remaining, ":")

			if index2 > 0 && isWord(remaining[0:index2]) {

				name := remaining[0:index2]
				unicode, ok := Emojis[name]

				if ok == true {

					element := NewElement("#text")
					element.SetLine(line)
					element.Text = unicode

					result = append(result, element)

				}

				remaining = remaining[index2+1:]

			} else {

				element := NewElement("#text")
				element.SetLine(line)
				element.Text = ":"

				result = append(result, element)

			}

		} else if is_media {

			remaining = remaining[2:]

			index1 := strings.Index(remaining, "](")

			if index1 >= 0 {

				index2 := index1 + strings.Index(remaining[index1+2:], ")")

				if index2 > index1 {

					text := remaining[0:index1]
					href := remaining[index1+2 : index2+2]

					if strings.HasSuffix(href, ".m4a") || strings.HasSuffix(href, ".mp3") || strings.HasSuffix(href, ".opus") {

						element := NewElement("audio")
						element.SetLine(line)
						element.SetAttribute("controls", "")
						element.SetAttribute("src", href)
						element.SetAttribute("title", text)

						result = append(result, element)
						remaining = remaining[index2+2+1:]

					} else if strings.HasSuffix(href, ".gif") || strings.HasSuffix(href, ".jpg") || strings.HasSuffix(href, ".png") {

						element := NewElement("img")
						element.SetLine(line)
						element.SetAttribute("alt", text)
						element.SetAttribute("src", href)

						result = append(result, element)
						remaining = remaining[index2+2+1:]

					} else {
						return result, fmt.Errorf("Unsupported media type \"%s\" at position %d", href, getTokenPosition(value, remaining))
					}

				} else {
					return result, fmt.Errorf("Expected token \"%s\" at positon %d", ")", getTokenPosition(value, remaining))
				}

			} else {
				return result, fmt.Errorf("Expected token \"%s\" at positon %d", "](", getTokenPosition(value, remaining))
			}

		} else if is_a {

			remaining = remaining[1:]

			index1 := strings.Index(remaining, "](")

			if index1 > 0 {

				index2 := index1 + strings.Index(remaining[index1+2:], ")")

				if index2 > index1 {

					text := remaining[0:index1]
					href := remaining[index1+2 : index2+2]

					element := NewElement("a")
					element.SetLine(line)
					element.SetText(text)
					element.SetAttribute("href", href)

					result = append(result, element)
					remaining = remaining[index2+2+1:]

				} else {
					return result, fmt.Errorf("Expected token \"%s\" at positon %d", ")", getTokenPosition(value, remaining))
				}

			} else {
				return result, fmt.Errorf("Expected token \"%s\" at positon %d", "](", getTokenPosition(value, remaining))
			}

		} else if is_footnote {

			remaining = remaining[2:]

			index2 := strings.Index(remaining, "]")

			if index2 > 0 {

				target := strings.ToLower(strings.TrimSpace(remaining[0:index2]))

				if target != "" {

					element := NewElement("a")
					element.SetLine(line)
					element.SetText("#" + target)
					element.SetAttribute("href", "#footnote-"+target)

					result = append(result, element)
					remaining = remaining[index2+1:]

				} else {
					return result, fmt.Errorf("Expected non-empty footnote at positon %d", getTokenPosition(value, remaining))
				}

			} else {
				return result, fmt.Errorf("Expected token \"%s\" at positon %d", "]", getTokenPosition(value, remaining))
			}

		} else if is_abbr {

			remaining = remaining[1:]

			index1 := strings.Index(remaining, "]{")

			if index1 >= 0 {

				index2 := index1 + strings.Index(remaining[index1+2:], "}")

				if index2 > index1 {

					text := remaining[0:index1]
					title := remaining[index1+2 : index2+2]

					element := NewElement("abbr")
					element.SetLine(line)
					element.SetText(text)
					element.SetAttribute("title", title)

					result = append(result, element)
					remaining = remaining[index2+2+1:]

				} else {
					return result, fmt.Errorf("Expected token \"%s\" at positon %d", "}", getTokenPosition(value, remaining))
				}

			} else {
				return result, fmt.Errorf("Expected token \"%s\" at positon %d", "]{", getTokenPosition(value, remaining))
			}

		} else if is_b {

			remaining = remaining[2:]

			if strings.Contains(remaining, "**") {

				index2 := strings.Index(remaining, "**")
				text := remaining[0:index2]

				element := NewElement("b")
				element.SetLine(line)
				element.SetText(text)

				result = append(result, element)
				remaining = remaining[index2+2:]

			} else {

				element := NewElement("#text")
				element.SetLine(line)
				element.SetText("**")

				result = append(result, element)

			}

		} else if is_code {

			remaining = remaining[1:]

			index2 := strings.Index(remaining, "`")

			if index2 > 0 {

				text := remaining[0:index2]

				element := NewElement("code")
				element.SetLine(line)
				element.SetText(text)

				result = append(result, element)
				remaining = remaining[index2+1:]

			} else {
				return result, fmt.Errorf("Expected token \"%s\" at positon %d", "`", getTokenPosition(value, remaining))
			}

		} else if is_em {

			remaining = remaining[1:]

			index2 := strings.Index(remaining, "*")

			if index2 > 0 {

				text := remaining[0:index2]

				element := NewElement("em")
				element.SetLine(line)
				element.SetText(text)

				result = append(result, element)
				remaining = remaining[index2+1:]

			} else {

				element := NewElement("#text")
				element.SetLine(line)
				element.SetText("*")

				result = append(result, element)

			}

		} else if is_del {

			remaining = remaining[1:]

			index2 := strings.Index(remaining, "~")

			if index2 > 0 {

				text := remaining[0:index2]

				element := NewElement("del")
				element.SetLine(line)
				element.SetText(text)

				result = append(result, element)
				remaining = remaining[index2+1:]

			} else {

				element := NewElement("#text")
				element.SetLine(line)
				element.SetText("~")

				result = append(result, element)

			}

		} else {

			seek_emoji := strings.Index(remaining, ":")
			seek_media := strings.Index(remaining, "![")
			seek_a_or_abbr_or_footnote := strings.Index(remaining, "[")
			seek_b := strings.Index(remaining, "**")
			seek_code := strings.Index(remaining, "`")
			seek_em := strings.Index(remaining, "*")
			seek_del := strings.Index(remaining, "~")

			seek := len(remaining)

			if seek_emoji != -1 && seek_emoji < seek {
				seek = seek_emoji
			}

			if seek_media != -1 && seek_media < seek {
				seek = seek_media
			}

			if seek_a_or_abbr_or_footnote != -1 && seek_a_or_abbr_or_footnote < seek {
				seek = seek_a_or_abbr_or_footnote
			}

			if seek_b != -1 && seek_b < seek {
				seek = seek_b
			}

			if seek_code != -1 && seek_code < seek {
				seek = seek_code
			}

			if seek_em != -1 && seek_em < seek {
				seek = seek_em
			}

			if seek_del != -1 && seek_del < seek {
				seek = seek_del
			}

			if seek > 0 {

				if len(result) > 0 {

					last_element := result[len(result)-1]

					if last_element.Type == "#text" {

						last_element.SetText(last_element.Text + remaining[0:seek])
						remaining = remaining[seek:]

					} else {

						element := NewElement("#text")
						element.SetLine(line)
						element.SetText(remaining[0:seek])

						result = append(result, element)
						remaining = remaining[seek:]

					}

				} else {

					element := NewElement("#text")
					element.SetLine(line)
					element.SetText(remaining[0:seek])

					result = append(result, element)
					remaining = remaining[seek:]

				}

			} else {
				break
			}

		}

	}

	if len(result) > 0 {

		for r := 0; r < len(result); r++ {

			element := result[r]

			if element.Type == "#text" {

				for _, abbreviation := range abbreviation_keys {

					index := strings.Index(element.Text, abbreviation)

					if index != -1 {

						description := Abbreviations[abbreviation]
						before := element.Text[0:index]
						after := element.Text[index+len(abbreviation):]

						element.SetText(before)

						abbr := NewElement("abbr")
						abbr.SetLine(line)
						abbr.SetText(abbreviation)
						abbr.SetAttribute("title", description)

						new_elements := []*Element{element, abbr}

						if after != "" {

							tmp := NewElement("#text")
							tmp.SetLine(line)
							tmp.SetText(after)

							new_elements = append(new_elements, tmp)

						}

						result = append(result[:r], append(new_elements, result[r+1:]...)...)

					}

				}

			}

		}

	}

	return result, nil

}
