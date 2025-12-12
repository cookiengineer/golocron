package markdown

import "strings"
import "testing"

func TestParse(t *testing.T) {

	t.Run("Parse(header)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description. It is not allowed to have newlines.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"### Example Title",
			"",
			"This is the first paragraph.",
			"",
		}, "\n"))

		document := Parse("/articles/example-title.md", bytes)

		if document.IsValid() == false {
			t.Errorf("Expected %v to be %v", document.IsValid(), true)
		}

		if len(document.Body) == 3 {

			if document.Body[0].Type != "h3" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "h3")
			}

			if document.Body[0].GetAttribute("id") != "example-title" {
				t.Errorf("Expected id %s to be %s", document.Body[0].GetAttribute("id"), "example-title")
			}

			if len(document.Body[0].Children) == 1 {

				if document.Body[0].Children[0].Type != "#text" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Type, "#text")
				}

				if document.Body[0].Children[0].Text != "Example Title" {
					t.Errorf("Expected \"%s\" to be \"%s\"", document.Body[0].Children[0].Text, "Example Title")
				}

			} else {
				t.Errorf("Expected %d elements to be %d", len(document.Body[0].Children), 1)
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type != "p" {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "p")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 3)
		}

	})

	t.Run("Parse(missing header)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"### Example Title",
			"",
			"This is the first paragraph.",
			"",
		}, "\n"))

		document := Parse("/articles/example-title.md", bytes)

		if document.IsValid() == true {
			t.Errorf("Expected %v to be %v", document.IsValid(), false)
		}

		if len(document.Body) == 3 {

			if document.Body[0].Type != "h3" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "h3")
			}

			if document.Body[0].GetAttribute("id") != "example-title" {
				t.Errorf("Expected id %s to be %s", document.Body[0].GetAttribute("id"), "example-title")
			}

			if len(document.Body[0].Children) == 1 {

				if document.Body[0].Children[0].Type != "#text" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Type, "#text")
				}

				if document.Body[0].Children[0].Text != "Example Title" {
					t.Errorf("Expected \"%s\" to be \"%s\"", document.Body[0].Children[0].Text, "Example Title")
				}

			} else {
				t.Errorf("Expected %d elements to be %d", len(document.Body[0].Children), 1)
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type != "p" {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "p")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 3)
		}

	})

	t.Run("Parse(inline elements)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- date: 2025-12-31",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- tags: software, network, privacy",
			"===",
			"",
			"### Example Title",
			"",
			"This is the first paragraph with `code`, some *emphasized* and **bold** ~and deleted~ text.",
			"",
		}, "\n"))

		document := Parse("/articles/example-title.md", bytes)

		if document.IsValid() != true {
			t.Errorf("Expected %v to be %v", document.IsValid(), true)
		}

		if len(document.Body) == 3 {

			if document.Body[0].Type != "h3" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "h3")
			}

			if document.Body[0].GetAttribute("id") != "example-title" {
				t.Errorf("Expected id %s to be %s", document.Body[0].GetAttribute("id"), "example-title")
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type == "p" {

				if len(document.Body[2].Children) == 9 {

					if document.Body[2].Children[0].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[0].Type, "#text")
					}

					if document.Body[2].Children[0].Text != "This is the first paragraph with" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[0].Text, "This is the first paragraph with")
					}

					if document.Body[2].Children[1].Type != "code" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[1].Type, "code")
					}

					if document.Body[2].Children[1].Text != "code" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[1].Text, "code")
					}

					if document.Body[2].Children[2].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[1].Type, "#text")
					}

					if document.Body[2].Children[2].Text != ", some" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[2].Text, ", some")
					}

					if document.Body[2].Children[3].Type != "em" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[3].Type, "em")
					}

					if document.Body[2].Children[3].Text != "emphasized" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[3].Text, "emphasized")
					}

					if document.Body[2].Children[4].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[4].Type, "#text")
					}

					if document.Body[2].Children[4].Text != "and" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[4].Text, "and")
					}

					if document.Body[2].Children[5].Type != "b" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[5].Type, "b")
					}

					if document.Body[2].Children[5].Text != "bold" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[5].Text, "bold")
					}

					if document.Body[2].Children[6].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[6].Type, "#text")
					}

					if document.Body[2].Children[6].Text != "" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[6].Text, "")
					}

					if document.Body[2].Children[7].Type != "del" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[7].Type, "del")
					}

					if document.Body[2].Children[7].Text != "and deleted" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[7].Text, "and deleted")
					}

					if document.Body[2].Children[8].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[8].Type, "#text")
					}

					if document.Body[2].Children[8].Text != "text." {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[8].Text, "text.")
					}

				} else {
					t.Errorf("Expected %d elements to be %d", len(document.Body[2].Children), 9)
				}

			} else {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "p")
			}


		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 3)
		}

	})

	t.Run("Parse(inline elements in headlines)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"### `Example` **Title**",
			"",
			"This is the first paragraph.",
			"",
		}, "\n"))

		document := Parse("/articles/example-title.md", bytes)

		if document.IsValid() != true {
			t.Errorf("Expected %v to be %v", document.IsValid(), true)
		}

		if len(document.Body) == 3 {

			if document.Body[0].Type == "h3" {

				if len(document.Body[0].Children) == 3 {

					if document.Body[0].Children[0].Type != "code" {
						t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Type, "code")
					}

					if document.Body[0].Children[0].Text != "Example" {
						t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Text, "Example")
					}

					if document.Body[0].Children[1].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[0].Children[1].Type, "#text")
					}

					if document.Body[0].Children[1].Text != "" {
						t.Errorf("Expected %s to be %s", document.Body[0].Children[1].Text, "")
					}

					if document.Body[0].Children[2].Type != "b" {
						t.Errorf("Expected %s to be %s", document.Body[0].Children[2].Type, "b")
					}

					if document.Body[0].Children[2].Text != "Title" {
						t.Errorf("Expected %s to be %s", document.Body[0].Children[2].Text, "Title")
					}

				} else {
					t.Errorf("Expected %d elements to be %d", len(document.Body[0].Children), 3)
				}

			} else {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "h3")
			}

			if document.Body[0].GetAttribute("id") != "example-title" {
				t.Errorf("Expected id %s to be %s", document.Body[0].GetAttribute("id"), "example-title")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 3)
		}

	})

	t.Run("Parse(inline emojis)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"This is the first paragraph under :construction: with :ghost: emojis.",
			"",
		}, "\n"))

		document := Parse("/articles/examples/example-title.md", bytes)

		if len(document.Body) == 1 {

			if document.Body[0].Type == "p" {

				if document.Body[0].Children[0].Type != "#text" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Type, "#text")
				}

				if document.Body[0].Children[0].Text != "This is the first paragraph under" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Text, "This is the first paragraph under")
				}

				if document.Body[0].Children[1].Type != "#text" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[1].Type, "#text")
				}

				if document.Body[0].Children[1].Text != " " + Emojis["construction"] + " " {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[1].Text, " " + Emojis["construction"] + " ")
				}

				if document.Body[0].Children[2].Type != "#text" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[2].Type, "#text")
				}

				if document.Body[0].Children[2].Text != "with" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[2].Text, "with")
				}

				if document.Body[0].Children[3].Type != "#text" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[3].Type, "#text")
				}

				if document.Body[0].Children[3].Text != " " + Emojis["ghost"] + " " {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[3].Text, " " + Emojis["ghost"] + " ")
				}

				if document.Body[0].Children[4].Type != "#text" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[4].Type, "#text")
				}

				if document.Body[0].Children[4].Text != "emojis." {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[4].Text, "emojis.")
				}

			} else {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "p")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 1)
		}

	})

	t.Run("Parse(images)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"![Relative Image 1](./path/to/image.png)",
			"![Relative Image 2](../path/to/image.png)",
			"![Relative Image 3](../../path/to/image.png)",
			"![Relative Image 4](../../../path/to/image.png)",
			"",
		}, "\n"))

		document := Parse("/articles/examples/example-title.md", bytes)

		if len(document.Body) == 4 {

			image1 := document.Body[0].RenderInto(document, "")

			if image1 != "<img alt=\"Relative Image 1\" src=\"https://cookie.engineer/articles/examples/path/to/image.png\"/>" {
				t.Errorf("Expected %s to resolve to %s", document.Body[0].GetAttribute("src"), "https://cookie.engineer/articles/examples/path/to/image.png")
			}

			image2 := document.Body[1].RenderInto(document, "")

			if image2 != "<img alt=\"Relative Image 2\" src=\"https://cookie.engineer/articles/path/to/image.png\"/>" {
				t.Errorf("Expected %s to resolve to %s", document.Body[1].GetAttribute("src"), "https://cookie.engineer/articles/path/to/image.png")
			}

			image3 := document.Body[2].RenderInto(document, "")

			if image3 != "<img alt=\"Relative Image 3\" src=\"https://cookie.engineer/path/to/image.png\"/>" {
				t.Errorf("Expected %s to resolve to %s", document.Body[2].GetAttribute("src"), "https://cookie.engineer/path/to/image.png")
			}

			image4 := document.Body[3].RenderInto(document, "")

			if image4 != "<img alt=\"Relative Image 4\" src=\"https://cookie.engineer/path/to/image.png\"/>" {
				t.Errorf("Expected %s to resolve to %s", document.Body[3].GetAttribute("src"), "https://cookie.engineer/path/to/image.png")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 4)
		}

	})

	t.Run("Parse(links)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"[Relative Link 1](./path/to/file.html)",
			"[Relative Link 2](../path/to/file.html)",
			"[Relative Link 3](../../path/to/file.html)",
			"[Relative Link 4](../../../path/to/file.html)",
			"",
		}, "\n"))

		document := Parse("/articles/examples/example-title.md", bytes)

		if len(document.Body) == 1 {

			if document.Body[0].Type == "p" {

				if len(document.Body[0].Children) == 4 {

					link1 := document.Body[0].Children[0].RenderInto(document, "")

					if link1 != "<a href=\"https://cookie.engineer/articles/examples/path/to/file.html\">Relative Link 1</a>" {
						t.Errorf("Expected %s to resolve to %s", document.Body[0].Children[0].GetAttribute("href"), "https://cookie.engineer/articles/examples/path/to/file.html")
					}

					link2 := document.Body[0].Children[1].RenderInto(document, "")

					if link2 != "<a href=\"https://cookie.engineer/articles/path/to/file.html\">Relative Link 2</a>" {
						t.Errorf("Expected %s to resolve to %s", document.Body[0].Children[1].GetAttribute("href"), "https://cookie.engineer/articles/path/to/file.html")
					}

					link3 := document.Body[0].Children[2].RenderInto(document, "")

					if link3 != "<a href=\"https://cookie.engineer/path/to/file.html\">Relative Link 3</a>" {
						t.Errorf("Expected %s to resolve to %s", document.Body[0].Children[2].GetAttribute("href"), "https://cookie.engineer/path/to/file.html")
					}

					link4 := document.Body[0].Children[3].RenderInto(document, "")

					if link4 != "<a href=\"https://cookie.engineer/path/to/file.html\">Relative Link 4</a>" {
						t.Errorf("Expected %s to resolve to %s", document.Body[0].Children[3].GetAttribute("href"), "https://cookie.engineer/path/to/file.html")
					}

				} else {
					t.Errorf("Expected %d children to be %d", len(document.Body[0].Children), 4)
				}

			} else {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "p")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 4)
		}

	})

	t.Run("Parse(lists)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"- This is a list item",
			"- This is another list item",
			"",
			"1. This is the first item",
			"2. This is the second item",
			"3. This is the third item",
			"",
			"+ This is a list item",
			"+ This is another list item",
			"",
		}, "\n"))

		document := Parse("/articles/examples/example-title.md", bytes)

		if document.IsValid() != true {
			t.Errorf("Expected %v to be %v", document.IsValid(), true)
		}

		if len(document.Body) == 5 {

			if document.Body[0].Type != "ul" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "ul")
			}

			if len(document.Body[0].Children) != 2 {
				t.Errorf("Expected %d children to be %d", len(document.Body[0].Children), 2)
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type != "ol" {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "ol")
			}

			if len(document.Body[2].Children) != 3 {
				t.Errorf("Expected %d children to be %d", len(document.Body[2].Children), 3)
			}

			if document.Body[3].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[3].Type, "#text")
			}

			if document.Body[4].Type != "ul" {
				t.Errorf("Expected %s to be %s", document.Body[4].Type, "ul")
			}

			if len(document.Body[4].Children) != 2 {
				t.Errorf("Expected %d children to be %d", len(document.Body[4].Children), 2)
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 5)
		}

	})

	t.Run("Parse(html elements)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"### `Example` **Title**",
			"",
			"<div id=\"my-id\" class=\"my-class\" data-name=\"something\">",
			"This is the first paragraph with `formatted` text.",
			"</div>",
			"",
			"- This is a list item",
			"- This is another list item",
			"",
		}, "\n"))

		document := Parse("/articles/example-title.md", bytes)

		if document.IsValid() != true {
			t.Errorf("Expected %v to be %v", document.IsValid(), true)
		}

		if len(document.Body) == 5 {

			if document.Body[0].Type != "h3" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "h3")
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type == "div" {

				if len(document.Body[2].Attributes) == 3 {

					if document.Body[2].GetAttribute("class") != "my-class" {
						t.Errorf("Expected attribute class value %s to be %s", document.Body[2].GetAttribute("class"), "my-class")
					}

					if document.Body[2].GetAttribute("data-name") != "something" {
						t.Errorf("Expected attribute data-name value %s to be %s", document.Body[2].GetAttribute("data-name"), "something")
					}

					if document.Body[2].GetAttribute("id") != "my-id" {
						t.Errorf("Expected attribute id value %s to be %s", document.Body[2].GetAttribute("id"), "my-id")
					}

				} else {
					t.Errorf("Expected %d attributes to be %d", len(document.Body[2].Attributes), 3)
				}

				if len(document.Body[2].Children) == 3 {

					if document.Body[2].Children[0].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[0].Type, "#text")
					}

					if document.Body[2].Children[0].Text != "This is the first paragraph with" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[0].Text, "This is the first paragraph with")
					}

					if document.Body[2].Children[1].Type != "code" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[1].Type, "code")
					}

					if document.Body[2].Children[1].Text != "formatted" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[1].Text, "formatted")
					}

					if document.Body[2].Children[2].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[1].Type, "#text")
					}

					if document.Body[2].Children[2].Text != "text." {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[2].Text, "text.")
					}

				} else {
					t.Errorf("Expected %d children to be %d", len(document.Body[2].Children), 3)
				}

			} else {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "div")
			}

			if document.Body[3].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[3].Type, "#text")
			}

			if document.Body[4].Type != "ul" {
				t.Errorf("Expected %s to be %s", document.Body[4].Type, "ul")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 5)
		}

	})

	t.Run("Parse(unclosed html elements)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"### `Example` **Title**",
			"",
			"<div id=\"my-id\" class=\"my-class\" data-name=\"something\">",
			"This is the first line.",
			"",
			"This is another line.",
			"",
			"- This is a list",
			"",
		}, "\n"))

		document := Parse("/articles/example-title.md", bytes)

		if document.IsValid() != true {
			t.Errorf("Expected %v to be %v", document.IsValid(), true)
		}

		if len(document.Body) == 7 {

			if document.Body[0].Type != "h3" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "h3")
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type == "div" {

				if len(document.Body[2].Attributes) == 3 {

					if document.Body[2].GetAttribute("class") != "my-class" {
						t.Errorf("Expected attribute class value %s to be %s", document.Body[2].GetAttribute("class"), "my-class")
					}

					if document.Body[2].GetAttribute("data-name") != "something" {
						t.Errorf("Expected attribute data-name value %s to be %s", document.Body[2].GetAttribute("data-name"), "something")
					}

					if document.Body[2].GetAttribute("id") != "my-id" {
						t.Errorf("Expected attribute id value %s to be %s", document.Body[2].GetAttribute("id"), "my-id")
					}

				} else {
					t.Errorf("Expected %d attributes to be %d", len(document.Body[2].Attributes), 3)
				}

				if len(document.Body[2].Children) == 1 {

					if document.Body[2].Children[0].Type != "#text" {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[0].Type, "#text")
					}

					if document.Body[2].Children[0].Text != "This is the first line." {
						t.Errorf("Expected %s to be %s", document.Body[2].Children[0].Text, "This is the first line.")
					}

				} else {
					t.Errorf("Expected %d children to be %d", len(document.Body[2].Children), 1)
				}

			} else {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "div")
			}

			if document.Body[3].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[3].Type, "#text")
			}

			if document.Body[4].Type != "p" {
				t.Errorf("Expected %s to be %s", document.Body[4].Type, "p")
			}

			if document.Body[5].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[5].Type, "#text")
			}

			if document.Body[6].Type != "ul" {
				t.Errorf("Expected %s to be %s", document.Body[6].Type, "ul")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 7)
		}

	})

	t.Run("Parse(selfclosed html elements)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description. It is not allowed to have newlines.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"### `Example` **Title**",
			"",
			"<img id=\"my-image\" src=\"/path/to/image.jpg\" width=\"1337\" height=\"137\"/>",
			"This is the first line.",
			"",
		}, "\n"))

		document := Parse("/articles/example-title.md", bytes)

		if document.IsValid() != true {
			t.Errorf("Expected %v to be %v", document.IsValid(), true)
		}

		if len(document.Body) == 4 {

			if document.Body[0].Type != "h3" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "h3")
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type == "img" {

				if len(document.Body[2].Attributes) == 4 {

					if document.Body[2].GetAttribute("id") != "my-image" {
						t.Errorf("Expected attribute id value %s to be %s", document.Body[2].GetAttribute("id"), "my-image")
					}

					if document.Body[2].GetAttribute("src") != "/path/to/image.jpg" {
						t.Errorf("Expected attribute src value %s to be %s", document.Body[2].GetAttribute("src"), "/path/to/image.jpg")
					}

					if document.Body[2].GetAttribute("width") != "1337" {
						t.Errorf("Expected attribute width value %s to be %s", document.Body[2].GetAttribute("width"), "1337")
					}

					if document.Body[2].GetAttribute("height") != "137" {
						t.Errorf("Expected attribute height value %s to be %s", document.Body[2].GetAttribute("height"), "137")
					}

				} else {
					t.Errorf("Expected %d attributes to be %d", len(document.Body[2].Attributes), 4)
				}

				if len(document.Body[2].Children) != 0 {
					t.Errorf("Expected %d children to be %d", len(document.Body[2].Children), 0)
				}

			} else {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "img")
			}

			if document.Body[3].Type != "p" {
				t.Errorf("Expected %s to be %s", document.Body[3].Type, "p")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 3)
		}

	})

	t.Run("Parse(inlined unsorted lists)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"<div id=\"my-id\" class=\"my-class\" data-name=\"something\">",
			"- This is a list item",
			"- This is another list item",
			"</div>",
			"",
			"This is a paragraph.",
			"",
		}, "\n"))

		document := Parse("/articles/examples/example-title.md", bytes)

		if len(document.Body) == 3 {

			if document.Body[0].Type != "div" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "div")
			}

			if len(document.Body[0].Children) == 1 {

				if document.Body[0].Children[0].Type != "ul" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Type, "ul")
				}

				if len(document.Body[0].Children[0].Children) != 2 {
					t.Errorf("Expected %d to be %d", len(document.Body[0].Children[0].Children), 2)
				}

			} else {
				t.Errorf("Expected %d children to be %d", len(document.Body[0].Children), 1)
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type != "p" {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "p")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 3)
		}

	})

	t.Run("Parse(inlined sorted lists)", func(t *testing.T) {

		bytes := []byte(strings.Join([]string{
			"===",
			"- title: Example Title",
			"- summary: This is an example article description.",
			"- date: 2025-12-31",
			"- tags: software, network, privacy",
			"===",
			"",
			"<div id=\"my-id\" class=\"my-class\" data-name=\"something\">",
			"1. This is the first item",
			"2. This is the second item",
			"3. This is the third item",
			"</div>",
			"",
			"This is a paragraph.",
			"",
		}, "\n"))

		document := Parse("/articles/examples/example-title.md", bytes)

		if len(document.Body) == 3 {

			if document.Body[0].Type != "div" {
				t.Errorf("Expected %s to be %s", document.Body[0].Type, "div")
			}

			if len(document.Body[0].Children) == 1 {

				if document.Body[0].Children[0].Type != "ol" {
					t.Errorf("Expected %s to be %s", document.Body[0].Children[0].Type, "ol")
				}

				if len(document.Body[0].Children[0].Children) != 3 {
					t.Errorf("Expected %d to be %d", len(document.Body[0].Children[0].Children), 3)
				}

			} else {
				t.Errorf("Expected %d children to be %d", len(document.Body[0].Children), 1)
			}

			if document.Body[1].Type != "#text" {
				t.Errorf("Expected %s to be %s", document.Body[1].Type, "#text")
			}

			if document.Body[2].Type != "p" {
				t.Errorf("Expected %s to be %s", document.Body[2].Type, "p")
			}

		} else {
			t.Errorf("Expected %d elements to be %d", len(document.Body), 3)
		}

	})
}
