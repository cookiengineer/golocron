
import { Editor } from "./Editor.mjs";

export const Init = (header, textarea, highlight, iframe, dialog) => {

	header    = typeof header === "object"    ? header    : null;
	textarea  = typeof textarea === "object"  ? textarea  : null;
	highlight = typeof highlight === "object" ? highlight : null;
	iframe    = typeof iframe === "object"    ? iframe    : null;
	dialog    = typeof dialog === "object"    ? dialog    : null;

	return new Promise((resolve, reject) => {

		let editor = new Editor(header, textarea, highlight, iframe, dialog);

		resolve(editor);

	});

};
