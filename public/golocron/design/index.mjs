
import { Init as InitEditor    } from "./editor/Init.mjs";
import { Init as InitHighlight } from "./highlight/Init.mjs";

(() => {

	let header    = document.querySelector("body > header");
	let textarea  = document.querySelector("body > div > aside > div > textarea");
	let highlight = document.querySelector("body > div > aside > div > pre");
	let iframe    = document.querySelector("body > div > main > iframe");
	let dialog    = document.querySelector("body > dialog");

	InitHighlight().then((hljs) => {

		InitEditor(header, textarea, highlight, iframe, dialog).then((editor) => {

			editor.Init();

			let query = window.location.search || "";
			if (query.startsWith("?file=")) {

				let file = query.substr(6);
				if (file.includes("&")) {
					file = file.substr(0, file.indexOf("&"));
				}

				editor.Open(file);

			}

			window.EDITOR = editor;

		});

	});

})();
