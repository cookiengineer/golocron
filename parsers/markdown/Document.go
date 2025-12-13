package markdown

import net_url "net/url"
import "regexp"
import "strings"
import "time"

type Document struct {
	URL  *net_url.URL `json:"url"`
	Meta struct {
		Author  string    `json:"author"`
		Title   string    `json:"title"`
		Summary string    `json:"summary"`
		Date    time.Time `json:"date"`
		Tags    []string  `json:"tags"`
		Image   string    `json:"image"`
	} `json:"meta"`
	Abbreviations map[string]string
	Statistics struct {
		Minutes int `json:"minutes"`
		Words   int `json:"words"`
	} `json:"statistics"`
	Body []*Element `json:"body"`
}

func (document *Document) AddElement(element *Element) {

	if element != nil && element.Type != "" {
		document.Body = append(document.Body, element)
	}

}

func (document *Document) Count() {

	var words int = 0

	for b := 0; b < len(document.Body); b++ {
		words += countWords(document.Body[b])
	}

	document.Statistics.Words = words
	document.Statistics.Minutes = int(document.Statistics.Words / 200)

}

func (document *Document) CloseElement() bool {

	var result bool = false
	var element *Element = nil

	if len(document.Body) > 0 {
		element = document.Body[len(document.Body)-1]
	}

	if element != nil {

		if element.Type != "#text" {
			document.Body = append(document.Body, NewElement("#text"))
		}

	}

	return result

}

func (document *Document) getLastElement() *Element {

	var element *Element = nil

	if len(document.Body) > 0 {
		element = document.Body[len(document.Body)-1]
	}

	return element

}

func (document *Document) IsValid() bool {

	var result bool

	if document.URL != nil && len(document.Body) > 0 {

		if document.Meta.Title != "" && document.Meta.Summary != "" {

			if document.Meta.Date.IsZero() == false && len(document.Meta.Tags) > 0 {
				result = true
			}

		}

	}

	return result

}

func (document *Document) ParseMeta(value string) {

	lines := strings.Split(strings.TrimSpace(value), "\n")

	if lines[0] == "===" {

		for l := 1; l < len(lines); l++ {

			line := strings.TrimSpace(lines[l])

			if strings.HasPrefix(line, "- ") && strings.Contains(line, ":") {

				key := strings.TrimSpace(line[2:strings.Index(line, ":")])
				val := strings.TrimSpace(line[strings.Index(line, ":")+1:])

				if key == "author" {

					document.SetAuthor(val)

				} else if key == "title" {

					document.SetTitle(val)

				} else if key == "summary" {

					document.SetSummary(val)

				} else if key == "date" {

					date, err := time.Parse("2006-01-02", val)

					if err == nil {
						document.SetDate(date)
					}

				} else if key == "tags" {

					values := strings.Split(val, ",")
					document.SetTags(values)

				} else if key == "image" {

					document.SetImage(val)

				}

			} else if line == "===" {
				break
			}

		}

	}

}

func (document *Document) ParseBody(value string) {

	lines := strings.Split(strings.TrimSpace(value), "\n")
	regexp_ul, _ := regexp.Compile("^([\\*\\-+]+)[[:space:]]")
	regexp_ol, _ := regexp.Compile("^([0-9]+)\\.[[:space:]]")

	first_line := 0

	if lines[0] == "===" {

		for l := 1; l < len(lines); l++ {

			line := strings.TrimSpace(lines[l])

			if line == "===" {
				first_line = l + 1
				break
			}

		}

	}

	for l := first_line; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])
		pointer := document.getLastElement()

		if pointer != nil && pointer.Type == "pre" {

			if line != "```" {
				// Preserve indention
				pointer.AddText(lines[l])
			} else {
				document.CloseElement()
			}

		} else if pointer != nil && pointer.Type == "table" {

			if strings.HasPrefix(line, "|") {
				pointer.AddText(line)
			} else {
				document.CloseElement()
			}

		} else if line == "" {

			if pointer == nil {
				// Do Nothing
			} else if pointer.Type == "pre" {
				// Do Nothing
			} else if pointer.Type == "div" || pointer.Type == "figure" || pointer.Type == "figcaption" || pointer.Type == "p" {
				document.CloseElement()
			} else if pointer.Type == "article" || pointer.Type == "dialog" || pointer.Type == "footer" || pointer.Type == "header" || pointer.Type == "main" {
				document.CloseElement()
			} else if pointer.Type == "ul" || pointer.Type == "ol" {
				document.CloseElement()
			} else if pointer.Type == "table" {
				document.CloseElement()
			} else if pointer != nil {
				document.CloseElement()
				pointer = nil
			}

		} else if strings.HasPrefix(line, "![") && strings.Contains(line, "](") && strings.HasSuffix(line, ")") {

			text := line[2:strings.Index(line, "](")]
			href := line[strings.Index(line, "](")+2 : strings.Index(line, ")")]

			element := NewElement("img")
			element.SetAttribute("alt", text)
			element.SetAttribute("src", href)

			if pointer != nil && pointer.IsBlockElement() == true {
				pointer.AddChild(element)
			} else {
				document.AddElement(element)
			}

		} else if len(line) > 3 && strings.HasPrefix(line, "```") {

			if pointer == nil || pointer.Type != "pre" {

				element := NewElement("pre")
				element.SetAttribute("class", strings.TrimSpace(line[3:]))

				document.AddElement(element)

			}

		} else if len(line) > 3 && strings.HasPrefix(line, "|") && strings.HasSuffix(line, "|") {

			if pointer == nil || pointer.Type != "table" {

				element := NewElement("table")
				element.SetText(line)

				document.AddElement(element)

			}

		} else if strings.HasPrefix(line, "####") {

			element := NewElement("h4")
			element.SetChildren(parseInlineElements(strings.TrimSpace(line[4:])))

			document.AddElement(element)

		} else if strings.HasPrefix(line, "###") {

			element := NewElement("h3")
			element.SetChildren(parseInlineElements(strings.TrimSpace(line[3:])))

			document.AddElement(element)

		} else if strings.HasPrefix(line, "##") {

			element := NewElement("h2")
			element.SetChildren(parseInlineElements(strings.TrimSpace(line[2:])))

			document.AddElement(element)

		} else if strings.HasPrefix(line, "#") {

			element := NewElement("h1")
			element.SetChildren(parseInlineElements(strings.TrimSpace(line[1:])))

			document.AddElement(element)

		} else if strings.HasPrefix(line, "<") && strings.HasSuffix(line, ">") {

			if strings.HasPrefix(line, "</") && strings.HasSuffix(line, ">") {

				element_type := line[2:len(line)-1]

				if element_type != "" {

					if pointer != nil && pointer.Type == element_type {
						pointer = nil
					}

				}

			} else if strings.HasPrefix(line, "<") && strings.HasSuffix(line, "/>") {

				element_type := ""

				if strings.Contains(line, " ") {
					element_type = line[1:strings.Index(line, " ")]
				} else {
					element_type = line[1:len(line)-2]
				}

				if element_type != "" {

					element := NewElement(element_type)

					// Don't expect a closing HTML tag
					element.is_block_element = false

					if strings.Contains(line, " ") {

						raw_attributes := strings.Split(line[strings.Index(line, " ")+1:len(line)-2], " ")

						for _, attribute := range raw_attributes {

							if strings.Contains(attribute, "=") {

								attribute_key := attribute[0:strings.Index(attribute, "=")]
								attribute_val := attribute[strings.Index(attribute, "=")+1:]

								if strings.HasPrefix(attribute_val, "\"") && strings.HasSuffix(attribute_val, "\"") {
									attribute_val = attribute_val[1:len(attribute_val)-1]
								} else if strings.HasPrefix(attribute_val, "'") && strings.HasSuffix(attribute_val, "'") {
									attribute_val = attribute_val[1:len(attribute_val)-1]
								}

								element.SetAttribute(attribute_key, attribute_val)

							} else if strings.TrimSpace(attribute) != "" {

								element.SetAttribute(strings.TrimSpace(attribute), "")

							}

						}

					}

					document.AddElement(element)

				}

			} else if strings.HasPrefix(line, "<") && strings.HasSuffix(line, ">") {

				element_type := ""

				if strings.Contains(line, " ") {
					element_type = line[1:strings.Index(line, " ")]
				} else {
					element_type = line[1:len(line)-1]
				}

				if element_type != "" {

					element := NewElement(element_type)

					// Expect a closing HTML tag
					element.is_block_element = true

					if strings.Contains(line, " ") {

						raw_attributes := strings.Split(line[strings.Index(line, " ")+1:len(line)-1], " ")

						for _, attribute := range raw_attributes {

							if strings.Contains(attribute, "=") {

								attribute_key := attribute[0:strings.Index(attribute, "=")]
								attribute_val := attribute[strings.Index(attribute, "=")+1:]

								if strings.HasPrefix(attribute_val, "\"") && strings.HasSuffix(attribute_val, "\"") {
									attribute_val = attribute_val[1:len(attribute_val)-1]
								} else if strings.HasPrefix(attribute_val, "'") && strings.HasSuffix(attribute_val, "'") {
									attribute_val = attribute_val[1:len(attribute_val)-1]
								}

								element.SetAttribute(attribute_key, attribute_val)

							} else if strings.TrimSpace(attribute) != "" {
								element.SetAttribute(strings.TrimSpace(attribute), "")
							}

						}

					}

					document.AddElement(element)

					pointer = document.getLastElement()

				}

			}

		} else if regexp_ul.MatchString(line) {

			item := NewElement("li")
			item.SetChildren(parseInlineElements(strings.TrimSpace(line[2:])))

			if pointer == nil {

				element := NewElement("ul")
				document.AddElement(element)
				pointer = document.getLastElement()

				pointer.AddChild(item)

			} else if pointer.Type == "ul" {

				pointer.AddChild(item)

			} else if pointer != nil && pointer.IsBlockElement() == true {

				inline_pointer := pointer.getLastChild()

				if inline_pointer == nil || inline_pointer.Type != "ul" {

					element := NewElement("ul")
					pointer.AddChild(element)
					inline_pointer = pointer.getLastChild()

				}

				inline_pointer.AddChild(item)

			} else {

				element := NewElement("ul")
				document.AddElement(element)
				pointer = document.getLastElement()

				pointer.AddChild(item)

			}

		} else if regexp_ol.MatchString(line) {

			item := NewElement("li")
			item.SetChildren(parseInlineElements(strings.TrimSpace(line[2:])))

			if pointer == nil {

				element := NewElement("ol")
				document.AddElement(element)
				pointer = document.getLastElement()

				pointer.AddChild(item)

			} else if pointer.Type == "ol" {

				pointer.AddChild(item)

			} else if pointer != nil && pointer.IsBlockElement() == true {

				inline_pointer := pointer.getLastChild()

				if inline_pointer == nil || inline_pointer.Type != "ol" {

					element := NewElement("ol")
					pointer.AddChild(element)
					inline_pointer = pointer.getLastChild()

				}

				inline_pointer.AddChild(item)

			} else {

				element := NewElement("ol")
				document.AddElement(element)
				pointer = document.getLastElement()

				pointer.AddChild(item)

			}

		} else if line != "" {

			if pointer == nil {

				element := NewElement("p")
				element.AddChildren(parseInlineElements(strings.TrimSpace(line)))

				document.AddElement(element)

			} else if pointer.Type == "p" {

				pointer.AddChildren(parseInlineElements(strings.TrimSpace(line)))

			} else if pointer != nil && pointer.IsBlockElement() == true {

				pointer.AddChildren(parseInlineElements(strings.TrimSpace(line)))

			} else {

				element := NewElement("p")
				element.AddChildren(parseInlineElements(strings.TrimSpace(line)))

				document.AddElement(element)

			}

		}

	}

}

func (document *Document) SetAuthor(value string) {
	document.Meta.Author = strings.TrimSpace(value)
}

func (document *Document) SetDate(value time.Time) bool {

	if value.IsZero() == false {

		document.Meta.Date = value

		return true

	}

	return false

}

func (document *Document) SetImage(value string) bool {

	tmp := strings.TrimSpace(value)

	if strings.HasPrefix(value, "./") || strings.HasPrefix(value, "/") {

		document.Meta.Image = tmp

		return true

	}

	return false

}

func (document *Document) SetSummary(value string) bool {

	tmp := strings.TrimSpace(value)

	if tmp != "" {

		document.Meta.Summary = tmp

		return true

	}

	return false

}

func (document *Document) SetTags(values []string) {

	filtered := make([]string, 0)

	for _, value := range values {

		tmp := strings.TrimSpace(strings.ToLower(value))

		if tmp != "" {
			filtered = append(filtered, tmp)
		}

	}

	document.Meta.Tags = filtered

}

func (document *Document) SetTitle(value string) {
	document.Meta.Title = strings.TrimSpace(value)
}

func (document *Document) Render(indent string) string {

	result := make([]string, 0)

	first_section_at := 0

	for e, element := range document.Body {

		if element.Type == "h3" {
			first_section_at = e
			break
		}

	}

	for e, element := range document.Body {

		if element.Type == "h1" {
			result = append(result, element.RenderInto(document, indent+"\t"))
		} else if element.Type == "h2" {
			result = append(result, element.RenderInto(document, indent+"\t"))
		} else if element.Type == "h3" {

			if first_section_at == e {
				result = append(result, indent+"<section>")
			} else {
				result = append(result, indent+"</section>")
				result = append(result, indent+"<section>")
			}

			result = append(result, element.RenderInto(document, indent+"\t"))

		} else if element.Type != "#text" {
			result = append(result, element.RenderInto(document, indent+"\t"))
		}

	}

	result = append(result, indent+"</section>")

	return strings.Join(result, "\n")

}
