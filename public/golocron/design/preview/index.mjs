
import { Init as InitHighlight } from "../highlight/Init.mjs";

(() => {

	InitHighlight().then((hljs) => {

		const themes = [ "sakura-vader", "sakura-dark", "sakura-light" ];

		let theme = sessionStorage.getItem("theme");
		if (theme !== null && typeof theme === "string") {
			document.body.setAttribute("data-theme", theme);
		} else {
			theme = "sakura-vader";
			document.body.setAttribute("data-theme", "sakura-vader");
		}

		document.addEventListener("keyup", (event) => {

			if ((event.ctrlKey || event.altKey) && event.key.toLowerCase() === "t") {

				let current_index = themes.indexOf(theme);
				let next_index    = (current_index + 1) % themes.length;

				theme = themes[next_index];
				document.body.setAttribute("data-theme", theme);
				sessionStorage.setItem("theme", theme);

			}

		});

		let codes = Array.from(document.querySelectorAll("pre[class]"));
		if (codes.length > 0) {

			codes.forEach((code) => {

				let language = code.className || "";
				if (language !== "") {
					hljs.highlightElement(code);
				}

			});

		}

	});

})();
