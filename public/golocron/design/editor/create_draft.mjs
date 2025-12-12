
import { datetime } from "./datetime.mjs";

export const create_draft = (url, title) => {

	if (url.pathname.startsWith("/projects/")) {

		let folder = url.pathname.substr(0, url.pathname.length - 3);
		let name   = folder.split("/").pop();
		let image  = folder + "/screenshot-01.png";

		return [
			"===",
			"- title: " + title,
			"- summary: A short summary",
			"- date: " + datetime(),
			"- tags: software, network, privacy",
			"- image: " + image,
			"===",
			"",
			"",
			"## " + title,
			"",
			"",
			"### Overview",
			"",
			"<div id=\"downloads\">",
			"[GitHub Repository](https://github.com/cookiengineer/" + name + ".git) | [GitLab Mirror](https://gitlab.com/cookiengineer/" + name + ".git) | [Download](" + folder + "/" + name + ".zip)",
			"</div>",
			"",
			"<figure id=\"screenshots\">",
			"![Screenshot showing ...](" + image + ")",
			"</figure>",
			"",
			"The `" + title + "` is a project that ...",
		].join("\n");

	} else {

		let folder = url.pathname.substr(0, url.pathname.length - 3);
		let image  = folder + "/teaser.jpg";

		return [
			"===",
			"- title: " + title,
			"- summary: A short summary",
			"- date: " + datetime(),
			"- tags: software, network, privacy",
			"- image: " + image,
			"===",
			"",
			"",
			"## " + title,
			"",
			"",
			"### Overview",
			"",
			"Lorem Ipsum Dolor Sit Amet...",
			"",
		].join("\n");

	}

};
