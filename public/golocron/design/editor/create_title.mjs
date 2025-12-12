
export const create_title = (url) => {

	let title = "New Article";

	if (url.pathname.startsWith("/") && url.pathname.endsWith(".md")) {

		let filename = url.pathname.split("/").pop();
		let words    = filename.substr(0, filename.length - 3).trim().split("-");

		if (words.length > 0) {

			let tmp = [];

			words.forEach((word) => {
				tmp.push(word.substr(0, 1).toUpperCase() + word.substr(1).toLowerCase());
			});

			title = tmp.join(" ");

		}

	}

	return title;

};
