package markdown

import "strings"

func stringTableBorder(lengths []int, alignment []string, border string) string {

	result := "|"

	for c, length := range lengths {

		switch alignment[c] {

		case "left":
			result += ":" + border[1:length+2] + "|"

		case "right":
			result += border[:length+1] + ":|"

		case "justify":
			result += ":" + border[1:length+1] + ":|"

		default:
			result += border[:length+2] + "|"

		}

	}

	result += "\n"

	return result

}

func stringTableRow(row []string, lengths []int, alignment []string, whitespace string) string {

	result := "|"

	for c, cell := range row {

		padding := lengths[c] - len(cell)

		switch alignment[c] {

		case "left":
			result += " " + cell + whitespace[:padding+1] + "|"

		case "right":
			result += whitespace[:padding+1] + cell + " |"

		case "justify", "center":
			left := padding / 2
			right := padding - left
			result += whitespace[:left+1] + cell + whitespace[:right+1] + "|"

		}

	}

	result += "\n"

	return result

}

func stringTable(element *Element) string {

	result := ""
	alignment := make([]string, 0)
	thead := make([]string, 0)
	tbody := make([][]string, 0)
	tfoot := make([][]string, 0)

	lines := strings.Split(strings.TrimSpace(element.Text), "\n")

	if len(lines) >= 2 {

		first_line := strings.Split(lines[0][1:len(lines[0])-1], "|")

		for _, field := range first_line {
			thead = append(thead, strings.TrimSpace(field))
		}

		if strings.Contains(lines[1], "---") {

			second_line := strings.Split(lines[1][1:len(lines[1])-1], "|")

			for _, align := range second_line {

				if strings.HasPrefix(align, ":-") && strings.HasSuffix(align, "-:") {
					alignment = append(alignment, "justify")
				} else if strings.HasPrefix(align, ":-") {
					alignment = append(alignment, "left")
				} else if strings.HasSuffix(align, "-:") {
					alignment = append(alignment, "right")
				} else {
					alignment = append(alignment, "center")
				}

			}

			if len(lines) > 2 {

				fill := "tbody"

				for l := 2; l < len(lines); l++ {

					line := lines[l]

					if strings.Contains(line, "---") {

						fill = "tfoot"

					} else {

						cells := strings.Split(line[1:len(line)-1], "|")
						row := make([]string, 0)

						for _, cell := range cells {
							row = append(row, strings.TrimSpace(cell))
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

					line := lines[l]

					if strings.Contains(line, "---") {

						fill = "tfoot"

					} else {

						cells := strings.Split(line[1:len(line)-1], "|")
						row := make([]string, 0)

						for _, cell := range cells {
							row = append(row, strings.TrimSpace(cell))
						}

						if fill == "tbody" {
							tbody = append(tbody, row)
						} else if fill == "tfoot" {
							tfoot = append(tfoot, row)
						}

					}

				}

			}

		}

	}

	if len(thead) > 0 || len(tbody) > 0 || len(tfoot) > 0 {

		lengths := make([]int, len(thead))

		for c, cell := range thead {
			lengths[c] = len(cell)
		}

		for _, row := range tbody {

			for c, cell := range row {

				if tmp := len(cell); tmp > lengths[c] {
					lengths[c] = tmp
				}

			}

		}

		for _, row := range tfoot {

			for c, cell := range row {

				if tmp := len(cell); tmp > lengths[c] {
					lengths[c] = tmp
				}

			}

		}

		border := ""
		whitespace := ""

		for _, length := range lengths {

			if length+2 > len(whitespace) {
				border = strings.Repeat("-", length + 2)
				whitespace = strings.Repeat(" ", length + 2)
			}

		}

		result += stringTableRow(thead, lengths, alignment, whitespace)
		result += stringTableBorder(lengths, alignment, border)

		if len(tbody) > 0 {

			for _, row := range tbody {
				result += stringTableRow(row, lengths, alignment, whitespace)
			}

		}

		if len(tfoot) > 0 {

			noalign := make([]string, len(alignment))
			result += stringTableBorder(lengths, noalign, border)

			for _, row := range tfoot {
				result += stringTableRow(row, lengths, alignment, whitespace)
			}

		}

	}

	return result

}
