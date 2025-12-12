
export const fix_highlighted_line = (line) => {

	let result     = line;
	let open_spans = 0;
	let matches    = line.match(/<\/?span[^>]*>/g) || [];

	for (let m = 0; m < matches.length; m++) {

		let tag = matches[m];
		if (tag.startsWith("</")) {
			open_spans--;
		} else {
			open_spans++;
		}

	}

	if (open_spans < 0) {

		// TODO: Need to add <span> in the front?

	} else if (open_spans > 0) {

		for (let o = 0; o < open_spans; o++) {
			result += "</span>";
		}

	}

	return result;

};

