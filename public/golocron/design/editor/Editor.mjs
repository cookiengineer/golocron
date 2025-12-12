
import { create_draft         } from "./create_draft.mjs";
import { create_title         } from "./create_title.mjs";
import { Dialog               } from "./Dialog.mjs";
import { datetime             } from "./datetime.mjs";
import { fix_highlighted_line } from "./fix_highlighted_line.mjs";
import hljs                     from "../highlight/highlight.mjs";

const base = new URL(".", import.meta.url);

export const Editor = function(header, textarea, highlight, iframe, dialog) {

	this.file = {
		url:    new URL("/weblog/articles/draft-" + datetime() + ".md", base),
		buffer: ""
	};

	this.file.buffer = create_draft(this.file.url, "New Weblog Article")

	this.dialog    = new Dialog(this, dialog);
	this.header    = header;
	this.highlight = highlight;
	this.textarea  = textarea;
	this.iframe    = iframe;

	this.__config = {
		base_url:     "http://localhost:3000",
		live_preview: true,
		documents:    {},
	};
	this.__listeners = {
		click:  null,
		input:  null,
		load:   null,
		scroll: null
	};

	this.Config(() => {
		this.Render();
	});

};

Editor.prototype = {

	FormatBlock: function(prefix, suffix) {

		let { before, selection, after } = this.SelectBlock();
		let was_formatted = false;

		if (selection.startsWith(prefix) && selection.endsWith(suffix)) {
			was_formatted = true;
		}

		let first_line = selection.split("\n").shift();
		let last_line  = selection.split("\n").pop();

		if (first_line.startsWith("```") && last_line.startsWith("```")) {

			let lines = selection.split("\n");
			if (lines.length >= 2) {
				was_formatted = true;
				selection = lines.slice(1, lines.length - 1).join("\n");
			}

		}

		if (was_formatted === true) {
			this.textarea.value          = before + selection + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + selection.length;
			this.textarea.focus();
		} else {
			this.textarea.value          = before + prefix + selection + suffix + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + prefix.length + selection.length + suffix.length;
			this.textarea.focus();
		}

	},

	FormatWord: function(prefix, suffix) {

		let { before, selection, after } = this.SelectWord();
		let was_formatted = false;

		if (selection.startsWith(prefix) && selection.endsWith(suffix)) {
			was_formatted = true;
		}

		if (selection.startsWith('**') && selection.endsWith('**')) {
			selection = selection.substr(2, selection.length - 4).trim();
		} else if (selection.startsWith('*') && selection.endsWith('*')) {
			selection = selection.substr(1, selection.length - 2).trim();
		} else if (selection.startsWith('`') && selection.endsWith('`')) {
			selection = selection.substr(1, selection.length - 2).trim();
		} else if (selection.startsWith('~') && selection.endsWith('~')) {
			selection = selection.substr(1, selection.length - 2).trim();
		}

		if (was_formatted === true) {
			this.textarea.value          = before + selection + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + selection.length;
			this.textarea.focus();
		} else {
			this.textarea.value          = before + prefix + selection + suffix + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + prefix.length + selection.length + suffix.length;
			this.textarea.focus();
		}

	},

	FormatList: function(prefix, suffix) {

		let { before, selection, after } = this.SelectLine();

		let regexp_ul     = /^([*\-+]+)\s/;
		let regexp_ol     = /^([0-9]+)\.\s/;
		let was_formatted = false;

		if (selection.startsWith(prefix) && selection.endsWith(suffix)) {
			was_formatted = true;
		}

		if (selection.startsWith('- [ ]') || selection.startsWith('- [x]')) {
			selection = selection.substr(5).trim();
		} else if (regexp_ul.test(selection)) {
			selection = selection.split(' ').slice(1).join(' ').trim();
		} else if (regexp_ol.test(selection)) {
			selection = selection.split(' ').slice(1).join(' ').trim();
		}

		if (was_formatted === true) {
			this.textarea.value          = before + selection + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + selection.length;
			this.textarea.focus();
		} else {
			this.textarea.value          = before + prefix + selection + suffix + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + prefix.length + selection.length + suffix.length;
			this.textarea.focus();
		}

	},

	FormatOrderedList: function() {

		let { before, selection, after } = this.SelectLine();

		let number        = 1;
		let regexp_ul     = /^([*\-+]+)\s/;
		let regexp_ol     = /^([0-9]+)\.\s/;
		let was_formatted = false;

		let { selection: block } = this.SelectBlock();

		if (block.includes("\n")) {

			let lines = block.split("\n");

			for (let l = 0; l < lines.length; l++) {

				let line = lines[l];
				if (line !== selection) {

					if (regexp_ol.test(line)) {

						let num = -1;

						try {
							num = Number.parseInt(line.split(".").shift(), 10);
						} catch (err) {
							num = -1;
						}

						if (num !== -1) {

							if (num >= number) {
								number = num + 1;
							}

						}

					}

				} else {
					break;
				}

			}

		}

		if (selection.startsWith('- [ ]') || selection.startsWith('- [x]')) {
			selection = selection.substr(5).trim();
		} else if (regexp_ul.test(selection)) {
			selection = selection.split(' ').slice(1).join(' ').trim();
		} else if (regexp_ol.test(selection)) {
			selection = selection.split(' ').slice(1).join(' ').trim();
			was_formatted = true;
		}

		let prefix = (number).toString() + ". ";
		let suffix = "";

		if (was_formatted === true) {
			this.textarea.value          = before + selection + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + selection.length;
			this.textarea.focus();
		} else {
			this.textarea.value          = before + prefix + selection + suffix + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + prefix.length + selection.length + suffix.length;
			this.textarea.focus();
		}

	},

	FormatHeadline: function(prefix, suffix) {

		let { before, selection, after } = this.SelectLine();
		let was_formatted = false;

		if (selection.startsWith(prefix) && selection.endsWith(suffix)) {
			was_formatted = true;
		}

		if (selection.startsWith("# ")) {
			selection = selection.substr(2).trim();
		} else if (selection.startsWith("## ")) {
			selection = selection.substr(3).trim();
		} else if (selection.startsWith("### ")) {
			selection = selection.substr(4).trim();
		} else if (selection.startsWith("#### ")) {
			selection = selection.substr(5).trim();
		} else if (selection.startsWith("##### ")) {
			selection = selection.substr(6).trim();
		}

		if (was_formatted === true) {
			this.textarea.value          = before + selection + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + selection.length;
			this.textarea.focus();
		} else {
			this.textarea.value          = before + prefix + selection + suffix + after;
			this.textarea.selectionStart = before.length;
			this.textarea.selectionEnd   = before.length + prefix.length + selection.length + suffix.length;
			this.textarea.focus();
		}

	},

	Config: function(callback) {

		callback = typeof callback === "function" ? callback : null;

		fetch("/golocron/api/config", {
			method:  "GET",
			headers: {
				"Accept": "application/json"
			}
		}).then((response) => {
			return response.json();
		}).then((config) => {

			this.__config = config;

			if (callback !== null) {
				callback(config);
			}

		});

	},

	Open: function(path) {

		let url = new URL(path, base);

		fetch("/golocron/api/open?file=" + url.pathname, {
			method:  "GET",
			headers: {
				"Accept": "text/markdown"
			}
		}).then((response) => {

			if (response.ok) {
				return response.text();
			} else {
				return create_draft(url, create_title(url));
			}

		}).then((buffer) => {

			this.file = {
				url:    url,
				buffer: buffer
			};

			this.Render();

		}).catch((err) => {
			console.error(err);
		});

	},

	Save: function(callback) {

		callback = typeof callback === "function" ? callback : null;

		fetch("/golocron/api/save", {
			method:  "POST",
			headers: {
				"Content-Type":  "text/markdown",
				"Content-Length": this.file.buffer.length,
				"X-Save-As":      this.file.url.pathname
			},
			body: this.file.buffer
		}).then((response) => {

			if (response.ok && response.status === 200) {

				let path_md   = this.file.url.pathname;
				let path_html = path_md.substr(0, path_md.length - 3) + ".html";

				if (this.iframe !== null) {
					this.iframe.setAttribute("src", path_html);
				}

				let save_button = this.header.querySelector("button[data-action=\"save\"]");
				if (save_button !== null) {
					save_button.setAttribute("disabled", "");
				}

				if (callback !== null) {
					callback(true);
				}

			} else {

				if (callback !== null) {
					callback(false);
				}

			}

		});

	},

	SelectBlock: function() {

		let prefix = this.textarea.value.substr(0, this.textarea.selectionStart);
		let suffix = this.textarea.value.substr(this.textarea.selectionEnd);
		let line   = this.textarea.value.substr(this.textarea.selectionStart, this.textarea.selectionEnd - this.textarea.selectionStart);

		if (prefix.indexOf('\n\n') !== -1) {

			let index = prefix.lastIndexOf('\n\n', this.textarea.selectionStart);

			line   = prefix.substr(index + 2) + line;
			prefix = prefix.substr(0, index + 2);

		} else {

			line   = prefix + line;
			prefix = '';

		}

		if (suffix.indexOf('\n\n') !== -1) {

			let index = suffix.indexOf('\n\n');

			line   = line + suffix.substr(0, index);
			suffix = suffix.substr(index);

		} else {

			line   = line + suffix;
			suffix = '';

		}

		return {
			before:    prefix,
			selection: line,
			after:     suffix
		};

	},

	SelectLine: function() {

		let prefix = this.textarea.value.substr(0, this.textarea.selectionStart);
		let suffix = this.textarea.value.substr(this.textarea.selectionEnd);
		let line   = this.textarea.value.substr(this.textarea.selectionStart, this.textarea.selectionEnd - this.textarea.selectionStart);

		if (prefix.indexOf('\n') !== -1) {

			let index = prefix.lastIndexOf('\n', this.textarea.selectionStart);

			line   = prefix.substr(index + 1) + line;
			prefix = prefix.substr(0, index + 1);

		} else {

			line   = prefix + line;
			prefix = '';

		}

		if (suffix.indexOf('\n') !== -1) {

			let index = suffix.indexOf('\n');

			line   = line + suffix.substr(0, index);
			suffix = suffix.substr(index);

		} else {

			line   = line + suffix;
			suffix = '';

		}

		return {
			before:    prefix,
			selection: line,
			after:     suffix
		};

	},

	SelectWord: function() {

		let prefix = this.textarea.value.substr(0, this.textarea.selectionStart);
		let suffix = this.textarea.value.substr(this.textarea.selectionEnd);
		let word   = this.textarea.value.substr(this.textarea.selectionStart, this.textarea.selectionEnd - this.textarea.selectionStart);

		if (prefix.indexOf(' ') !== -1) {

			let index = prefix.lastIndexOf(' ', this.textarea.selectionStart);
			let limit = prefix.lastIndexOf('\n');
			if (limit !== -1 && index < limit) {
				index = limit;
			}

			word   = prefix.substr(index + 1) + word;
			prefix = prefix.substr(0, index + 1);

		} else {

			let limit = prefix.lastIndexOf('\n');
			if (limit !== -1) {
				word   = prefix.substr(limit + 1) + word;
				prefix = prefix.substr(0, limit + 1);
			} else {
				word   = prefix + word;
				prefix = '';
			}

		}

		if (suffix.indexOf(' ') !== -1) {

			let index = suffix.indexOf(' ');
			let limit = suffix.indexOf('\n');
			if (limit !== -1 && index > limit) {
				index = limit;
			}

			word   = word + suffix.substr(0, index);
			suffix = suffix.substr(index);

		} else {

			let limit = suffix.indexOf('\n');
			if (limit !== -1) {
				word   = word + suffix.substr(0, limit);
				suffix = suffix.substr(limit);
			} else {
				word   = word + suffix;
				suffix = '';
			}

		}

		return {
			before:    prefix,
			selection: word,
			after:     suffix
		};

	},

	Init: function() {

		if (this.__listeners.click === null) {

			this.__listeners.click = (event) => {

				let target = event.target;
				let tagname = target.tagName.toLowerCase();

				if (tagname === "svg") {
					target  = target.parentNode;
					tagname = target.tagName.toLowerCase();
				} else if (tagname === "circle" || tagname === "path" || tagname === "rect") {
					target  = target.parentNode.parentNode;
					tagname = target.tagName.toLowerCase();
				} else if (tagname === "button") {
					// Do Nothing
				}

				if (tagname === "button") {

					let action = target.getAttribute("data-action");
					let format = target.getAttribute("data-format");

					if (action === "open") {
						this.dialog.Show();
					} else if (action === "save") {
						this.Save();
					} else if (format === "h1") {
						this.FormatHeadline("# ", "");
					} else if (format === "h2") {
						this.FormatHeadline("## ", "");
					} else if (format === "h3") {
						this.FormatHeadline("### ", "");
					} else if (format === "h4") {
						this.FormatHeadline("#### ", "");
					} else if (format === "b") {
						this.FormatWord("**", "**");
					} else if (format === "i") {
						this.FormatWord("*", "*");
					} else if (format === "del") {
						this.FormatWord("~", "~");
					} else if (format === "q") {
						this.FormatWord("`", "`");
					} else if (format === "pre") {

						let language = window.prompt("Enter programming language:");
						if (language !== null && language !== "") {
							this.FormatBlock("```" + language + "\n", "\n```");
						} else {
							this.FormatBlock("```\n", "\n```");
						}

					} else if (format === "a") {

						let url = window.prompt("Enter URL:");
						if (url !== null && url !== "") {
							this.FormatWord("[", "](" + url + ")");
						} else {
							this.FormatWord("[", "](...)");
						}

					} else if (format === "img") {

						let url = window.prompt("Enter URL:");
						if (url !== null && url !== "") {
							this.FormatWord("![", "](" + url + ")");
						} else {
							this.FormatWord("![", "](...)");
						}

					} else if (format === "ul") {
						this.FormatList("- ", "");
					} else if (format === "ol") {
						this.FormatOrderedList();
					} else if (format === "dl") {
						this.FormatList("- [ ] ", "");
					}

				}

			};

			this.__listeners.load = () => {

				try {

					let doc = this.iframe.contentDocument || this.iframe.contentWindow.document;
					if (doc.body.children.length === 0) {
						doc.body.style.cssText = [
							"display: flex;",
							"justify-content: center;",
							"align-items: center;",
							"height: 100vh;",
							"color: #cc0000;",
							"background-color: #202020;",
							"font-size: 25px;",
							"text-align:center;",
							"overflow: hidden;"
						].join("\n");
					}

				} catch (err) {
					console.error("Cannot access iframe Content: " + err);
				}

			};

			this.__listeners.input = () => {

				if (this.file.buffer !== this.textarea.value) {

					this.file.buffer = this.textarea.value;

					let save_button = this.header.querySelector("button[data-action=\"save\"]");
					if (save_button !== null) {
						save_button.removeAttribute("disabled");
					}

				}

				this.RenderHighlight();

			};

			this.__listeners.scroll = () => {

				let boundary = this.highlight.clientHeight - this.textarea.clientHeight;

				if (this.textarea.scrollHeight - this.textarea.scrollTop <= this.textarea.clientHeight + boundary) {

					// XXX: The textarea can overflow, Browser behavior is different from other elements, apparently
					this.textarea.scrollTop   = this.textarea.scrollTop - boundary;
					this.highlight.scrollTop  = this.textarea.scrollTop;
					this.highlight.scrollLeft = this.textarea.scrollLeft;

				} else {
					this.highlight.scrollTop  = this.textarea.scrollTop;
					this.highlight.scrollLeft = this.textarea.scrollLeft;
				}

			};

			this.textarea.addEventListener("input", this.__listeners.input, true);
			this.textarea.addEventListener("scroll", this.__listeners.scroll, true);

			this.header.addEventListener("click", this.__listeners.click, true);
			this.iframe.addEventListener("load", this.__listeners.load, true);

		}

	},

	Render: function() {

		let path_md   = this.file.url.pathname;
		let path_html = path_md.substr(0, path_md.length - 3) + ".html";

		let label = this.header.querySelector("label");
		if (label !== null) {
			label.innerHTML = "<a href=\"" + path_html + "\" target=\"_blank\">" + path_md + "</a>";
		}

		if (this.textarea !== null) {
			this.textarea.value = this.file.buffer;
		}

		if (this.iframe !== null) {

			if (this.__config.live_preview === true) {
				this.iframe.setAttribute("src", path_html);
			} else {
				this.iframe.setAttribute("src", "/golocron/api/render?file=" + this.file.url.pathname);
			}

		}

		if (this.highlight !== null) {
			this.RenderHighlight();
		}

		let save_button = this.header.querySelector("button[data-action=\"save\"]");
		if (save_button !== null) {
			save_button.setAttribute("disabled", "");
		}

	},

	RenderHighlight: function() {

		let { value: highlighted } = hljs.highlight(this.textarea.value, {
			language: "markdown"
		});

		let lines_markdown    = this.textarea.value.split("\n");
		let lines_highlighted = highlighted.split(/\r?\n/);

		let is_header       = false;
		let inline_language = null;
		let mapped_lines    = lines_markdown.map((line, l) => {

			if (line === "===") {

				if (is_header === false) {
					is_header = true;
				} else {
					is_header = false;
				}

				return "<span><span class=\"hljs-section\">" + (lines_markdown[l] || '') + "</span></span>";

			} else if (is_header === true) {

				if (line.startsWith("- ") && line.includes(": ")) {

					let tmp1 = line.substr(0, 2);
					let tmp2 = line.substr(2, line.indexOf(": ") - 2);
					let tmp3 = line.substr(line.indexOf(": "), 2);
					let tmp4 = line.substr(line.indexOf(": ") + 2);

					return "<span>" + tmp1 + "<span class=\"hljs-attr\">" + tmp2 + "</span>" + tmp3 + "<span class=\"hljs-string\">" + tmp4 + "</span></span>";

				} else {
					return "<span><span class=\"hljs-del\">" + (lines_markdown[l] || '') + "</span></span>";
				}

			} else if (line.startsWith("```")) {

				if (inline_language === null) {
					inline_language = line.substr(3).trim();
				} else {
					inline_language = null;
				}

				return "<span><span class=\"hljs-section\">" + (lines_markdown[l] || '') + "</span></span>";

			} else if (inline_language !== null) {

				let { value: inline_highlighted } = hljs.highlight(lines_markdown[l], {
					language: inline_language
				});

				return "<span>" + fix_highlighted_line(inline_highlighted) + "</span>";

			} else {

				return "<span>" + (fix_highlighted_line(lines_highlighted[l]) || '') + "</span>";

			}

		}).join("");

		this.highlight.innerHTML = mapped_lines;

	},

};
