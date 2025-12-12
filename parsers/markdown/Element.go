package markdown

import "html"
import "sort"
import "strings"

type Element struct {
	Type             string            `json:"type"`
	Text             string            `json:"text"`
	Attributes       map[string]string `json:"attributes"`
	Children         []*Element        `json:"children"`
	is_block_element bool              `json:"-"`
}

func NewElement(typ string) *Element {

	var element Element

	element.Type = strings.TrimSpace(typ)
	element.Attributes = make(map[string]string)
	element.Children = make([]*Element, 0)
	element.is_block_element = false

	// There's no clean place to do this, as it's part of the CSS specification
	switch typ {
	case "div", "figure", "figcaption", "p":
		element.is_block_element = true
	case "pre":
		element.is_block_element = true
	case "article", "dialog", "footer", "header", "main":
		element.is_block_element = true
	default:
		element.is_block_element = false
	}

	return &element

}

func (element *Element) AddChild(child *Element) {

	element.Children = append(element.Children, child)

	if element.Type == "h1" || element.Type == "h2" || element.Type == "h3" || element.Type == "h4" {
		element.Attributes["id"] = generateId(element)
	}

}

func (element *Element) AddChildren(children []*Element) {

	filtered := make([]*Element, 0)

	for _, child := range children {

		if child != nil {
			filtered = append(filtered, child)
		}

	}

	for _, child := range filtered {
		element.Children = append(element.Children, child)
	}

	if element.Type == "h1" || element.Type == "h2" || element.Type == "h3" || element.Type == "h4" {
		element.Attributes["id"] = generateId(element)
	}

}

func (element *Element) AddText(value string) {

	if element.Text != "" {
		element.Text = element.Text + "\n" + value
	} else {
		element.Text = value
	}

	if element.Type == "h1" || element.Type == "h2" || element.Type == "h3" || element.Type == "h4" {
		element.Attributes["id"] = generateId(element)
	}

}

func (element *Element) getAttributes() []string {

	result := make([]string, 0)

	for name , _ := range element.Attributes {
		result = append(result, name)
	}

	sort.Strings(result)

	return result

}

func (element *Element) GetAttribute(key string) string {

	var result string

	value, ok := element.Attributes[key]

	if ok == true {
		result = value
	}

	return result

}

func (element *Element) getLastChild() *Element {

	var child *Element = nil

	if len(element.Children) > 0 {
		child = element.Children[len(element.Children)-1]
	}

	return child

}

func (element *Element) HasAttribute(key string) bool {

	var result bool

	_, ok := element.Attributes[key]

	if ok == true {
		result = true
	}

	return result

}

func (element *Element) IsBlockElement() bool {
	return element.is_block_element
}

func (element *Element) RenderInto(document *Document, indent string) string {

	result := ""

	if element.Type == "b" || element.Type == "em" || element.Type == "del" {

		result += indent + "<" + element.Type + ">" + element.Text + "</" + element.Type + ">"

	} else if element.Type == "code" {

		result += indent + "<" + element.Type + ">" + html.EscapeString(element.Text) + "</" + element.Type + ">"

	} else if element.Type == "a" {

		result += indent + "<" + element.Type

		if len(element.Attributes) > 0 {

			attribute_names := element.getAttributes()
			is_same_site := false

			for _, name := range attribute_names {

				value := element.GetAttribute(name)

				if name == "href" {
					value = resolveURL(document.URL, value)
					is_same_site = strings.HasPrefix(value, "https://cookie.engineer/")
				} else if name == "ping" {
					value = resolveURL(document.URL, value)
				}

				if value != "" {
					result += " " + name + "=\"" + value + "\""
				} else {
					result += " " + name
				}

			}

			if is_same_site == false && element.HasAttribute("target") == false {
				result += " target=\"_blank\""
			}

		}

		result += ">"
		result += element.Text
		result += "</" + element.Type + ">"

	} else if element.Type == "abbr" {

		result += indent + "<" + element.Type

		if len(element.Attributes) > 0 {

			attribute_names := element.getAttributes()

			for _, name := range attribute_names {

				value := element.GetAttribute(name)

				if value != "" {
					result += " " + name + "=\"" + value + "\""
				} else {
					result += " " + name
				}

			}

		}

		result += ">"
		result += element.Text
		result += "</" + element.Type + ">"

	} else if element.Type == "img" || element.Type == "audio" || element.Type == "video" {

		result += indent + "<" + element.Type

		if len(element.Attributes) > 0 {

			attribute_names := element.getAttributes()

			for _, name := range attribute_names {

				value := element.GetAttribute(name)

				if name == "src" {
					value = resolveURL(document.URL, value)
				}

				if value != "" {
					result += " " + name + "=\"" + value + "\""
				} else {
					result += " " + name
				}

			}

		}

		result += "/>"

	} else if element.Type == "h1" || element.Type == "h2" || element.Type == "h3" || element.Type == "h4" {

		id := element.GetAttribute("id")

		if id == "" {
			id = generateId(element)
		}

		result += indent + "<" + element.Type + " id=\"" + id + "\">"

		for c, child := range element.Children {

			result += child.RenderInto(document, "")

			if c < len(element.Children) - 1 {
				result += " "
			}

		}

		result += "</" + element.Type + ">"

	} else if element.Type == "div" || element.Type == "figure" || element.Type == "figcaption" || element.Type == "p" {

		if len(element.Children) > 1 {

			result += indent + "<" + element.Type

			attribute_names := element.getAttributes()

			for _, name := range attribute_names {

				value := element.GetAttribute(name)

				if value != "" {
					result += " " + name + "=\"" + value + "\""
				} else {
					result += " " + name
				}

			}

			result += ">\n"

			for _, child := range element.Children {
				result += child.RenderInto(document, indent+"\t") + "\n"
			}

			result += indent + "</" + element.Type + ">"

		} else if len(element.Children) == 1 {

			result += indent + "<" + element.Type

			attribute_names := element.getAttributes()

			for _, name := range attribute_names {

				value := element.GetAttribute(name)

				if value != "" {
					result += " " + name + "=\"" + value + "\""
				} else {
					result += " " + name
				}

			}

			result += ">" + element.Children[0].RenderInto(document, "") + "</" + element.Type + ">"

		}

	} else if element.Type == "article" || element.Type == "dialog" || element.Type == "footer" || element.Type == "header" || element.Type == "main" {

		result += indent + "<" + element.Type

		attribute_names := element.getAttributes()

		for _, name := range attribute_names {

			value := element.GetAttribute(name)

			if value != "" {
				result += " " + name + "=\"" + value + "\""
			} else {
				result += " " + name
			}

		}

		result += ">\n"


		for _, child := range element.Children {
			result += child.RenderInto(document, indent+"\t") + "\n"
		}

		result += indent + "</" + element.Type + ">"

	} else if element.Type == "pre" {

		class, ok := element.Attributes["class"]

		text := html.EscapeString(element.Text)

		if ok == true {
			result += indent + "<pre class=\"" + class + "\">\n"
			result += text + "\n"
			result += indent + "</pre>"
		} else {
			result += indent + "<pre>\n"
			result += text + "\n"
			result += indent + "</pre>"
		}

	} else if element.Type == "table" {

		alignment := make([]string, 0)
		thead := make([]string, 0)
		tbody := make([][]string, 0)
		tfoot := make([][]string, 0)

		lines := strings.Split(strings.TrimSpace(element.Text), "\n")

		if len(lines) >= 2 {

			first_line := strings.Split(lines[0][1:len(lines[0])-1], "|")

			for f := 0; f < len(first_line); f++ {
				thead = append(thead, strings.TrimSpace(first_line[f]))
			}

			if strings.Contains(lines[1], "---") {

				second_line := strings.Split(lines[1][1:len(lines[1])-1], "|")

				for s := 0; s < len(second_line); s++ {

					cell := second_line[s]

					if strings.HasPrefix(cell, ":-") && strings.HasSuffix(cell, "-:") {
						alignment = append(alignment, "justify")
					} else if strings.HasPrefix(cell, ":-") {
						alignment = append(alignment, "left")
					} else if strings.HasSuffix(cell, "-:") {
						alignment = append(alignment, "right")
					} else {
						alignment = append(alignment, "center")
					}

				}

				if len(lines) > 2 {

					fill := "tbody"

					for l := 2; l < len(lines); l++ {

						if strings.Contains(lines[l], "---") {
							fill = "tfoot"
						} else {

							cells := strings.Split(lines[l][1:len(lines[l])-1], "|")
							row := make([]string, 0)

							for c := 0; c < len(cells); c++ {
								row = append(row, strings.TrimSpace(cells[c]))
							}

							if fill == "tbody" {
								tbody = append(tbody, row)
							} else if fill == "tfoot" {
								tfoot = append(tfoot, row)
							}

						}

					}

				}

			} else {

				for t := 0; t < len(thead); t++ {
					alignment = append(alignment, "center")
				}

				if len(lines) > 1 {

					fill := "tbody"

					for l := 1; l < len(lines); l++ {

						if strings.Contains(lines[l], "---") {
							fill = "tfoot"
						} else {

							cells := strings.Split(lines[l][1:len(lines[l])-1], "|")
							row := make([]string, 0)

							for c := 0; c < len(cells); c++ {
								row = append(row, strings.TrimSpace(cells[c]))
							}

							if fill == "tbody" {
								tbody = append(tbody, row)
							} else if fill == "tfoot" {
								tfoot = append(tbody, row)
							}

						}

					}

				}

			}

		}

		if len(thead) > 0 || len(tbody) > 0 || len(tfoot) > 0 {

			result += indent + "<" + element.Type + ">\n"

			if len(thead) > 0 {

				result += indent + "\t<thead>\n"
				result += indent + "\t\t<tr>\n"

				for t := 0; t < len(thead); t++ {

					elements := parseInlineElements(strings.TrimSpace(thead[t]))
					align := alignment[t]

					if align == "center" {
						result += indent + "\t\t\t<th align=\"center\">"
					} else if align == "left" {
						result += indent + "\t\t\t<th align=\"left\">"
					} else if align == "right" {
						result += indent + "\t\t\t<th align=\"right\">"
					} else if align == "justify" {
						result += indent + "\t\t\t<th align=\"justify\">"
					}

					for e := 0; e < len(elements); e++ {
						result += elements[e].RenderInto(document, "")
					}

					result += "</th>\n"

				}

				result += indent + "\t\t</tr>\n"
				result += indent + "\t</thead>\n"

			}

			if len(tbody) > 0 {

				result += indent + "\t<tbody>\n"

				for t := 0; t < len(tbody); t++ {

					result += indent + "\t\t<tr>\n"

					cells := tbody[t]

					for c := 0; c < len(cells); c++ {

						elements := parseInlineElements(strings.TrimSpace(cells[c]))
						align := alignment[c]

						if align == "center" {
							result += indent + "\t\t\t<td align=\"center\">"
						} else if align == "left" {
							result += indent + "\t\t\t<td align=\"left\">"
						} else if align == "right" {
							result += indent + "\t\t\t<td align=\"right\">"
						} else if align == "justify" {
							result += indent + "\t\t\t<td align=\"justify\">"
						}

						for e := 0; e < len(elements); e++ {
							result += elements[e].RenderInto(document, "")
						}

						result += "</td>\n"

					}

					result += indent + "\t\t</tr>\n"

				}

				result += indent + "\t</tbody>\n"

			}

			if len(tfoot) > 0 {

				result += indent + "\t<tfoot>\n"

				for t := 0; t < len(tfoot); t++ {

					cells := tfoot[t]

					result += indent + "\t\t<tr>\n"

					for c := 0; c < len(cells); c++ {

						elements := parseInlineElements(strings.TrimSpace(cells[c]))
						align := alignment[c]

						if align == "center" {
							result += indent + "\t\t\t<td align=\"center\">"
						} else if align == "left" {
							result += indent + "\t\t\t<td align=\"left\">"
						} else if align == "right" {
							result += indent + "\t\t\t<td align=\"right\">"
						} else if align == "justify" {
							result += indent + "\t\t\t<td align=\"justify\">"
						}

						for e := 0; e < len(elements); e++ {
							result += elements[e].RenderInto(document, "")
						}

						result += "</td>\n"

					}

					result += indent + "\t\t</tr>\n"

				}

				result += indent + "\t</tfoot>\n"

			}

			result += indent + "</" + element.Type + ">"

		}

	} else if element.Type == "ol" || element.Type == "ul" {

		result += indent + "<" + element.Type + ">\n"

		for c := 0; c < len(element.Children); c++ {
			result += element.Children[c].RenderInto(document, indent+"\t") + "\n"
		}

		result += indent + "</" + element.Type + ">"

	} else if element.Type == "li" {

		result += indent + "<" + element.Type + ">"

		for c, child := range element.Children {

			result += child.RenderInto(document, "")

			if c < len(element.Children) - 1 {
				result += " "
			}

		}

		result += "</" + element.Type + ">"

	} else if element.Type == "#text" {

		result += indent + element.Text

	} else if element.Type != "" {

		// TODO: Should we allow unsupported HTML tags? Potential XSS?

	}

	return result

}

func (element *Element) SetAttribute(key string, value string) {

	if element.Type == "a" && key == "href" {

		if strings.HasPrefix(value, "https://github.com") {
			element.Attributes["class"] = "icon-github"
		} else if strings.HasPrefix(value, "https://gitlab.com") {
			element.Attributes["class"] = "icon-gitlab"
		} else if strings.HasPrefix(value, "https://instagram.com") {
			element.Attributes["class"] = "icon-instagram"
		} else if strings.HasPrefix(value, "https://linkedin.com") {
			element.Attributes["class"] = "icon-linkedin"
		} else if strings.HasPrefix(value, "https://medium.com") {
			element.Attributes["class"] = "icon-medium"
		} else if strings.HasPrefix(value, "https://reddit.com") {
			element.Attributes["class"] = "icon-reddit"
		} else if strings.HasPrefix(value, "https://old.reddit.com") {
			element.Attributes["class"] = "icon-reddit"
		} else if strings.HasPrefix(value, "https://cookie.engineer") {
			// Do Nothing
		} else if strings.HasPrefix(value, "../") {
			// Do Nothing
		} else if strings.HasPrefix(value, "./") {
			// Do Nothing
		} else if strings.HasPrefix(value, "/") {
			// Do Nothing
		} else if strings.HasPrefix(value, "https://") || strings.HasPrefix(value, "http://") {
			element.Attributes["class"] = "icon-website"
		}

		if strings.HasSuffix(value, ".go") ||
			strings.HasSuffix(value, ".js") ||
			strings.HasSuffix(value, ".mjs") ||
			strings.HasSuffix(value, ".pdf") ||
			strings.HasSuffix(value, ".tar.gz") ||
			strings.HasSuffix(value, ".tar.xz") ||
			strings.HasSuffix(value, ".zip") {
			element.Attributes["class"] = "icon-download"
		} else if strings.HasPrefix(value, "#") {
			element.Attributes["class"] = "icon-section"
		}

		element.Attributes[key] = strings.TrimSpace(value)

	} else if element.Type == "img" && key == "src" {

		element.Attributes[key] = strings.TrimSpace(value)

	} else {

		element.Attributes[key] = strings.TrimSpace(value)

	}

}

func (element *Element) SetChildren(children []*Element) {

	filtered := make([]*Element, 0)

	for _, child := range children {

		if child != nil {
			filtered = append(filtered, child)
		}

	}

	element.Children = filtered

	if element.Type == "h1" || element.Type == "h2" || element.Type == "h3" || element.Type == "h4" {
		element.Attributes["id"] = generateId(element)
	}

}

func (element *Element) SetClass(value string) {
	element.Attributes["class"] = strings.TrimSpace(value)
}

func (element *Element) SetText(value string) {

	element.Text = strings.TrimSpace(value)

	if element.Type == "h1" || element.Type == "h2" || element.Type == "h3" || element.Type == "h4" {
		element.Attributes["id"] = generateId(element)
	}

}

