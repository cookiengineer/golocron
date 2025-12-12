
const is_valid_filepath = (path) => {

	// TODO: This should validate for existing folders somehow
	if (path.startsWith("/") && !path.includes("..") && path.endsWith(".md")) {
		return true;
	} else {
		return false;
	}

};

export const Dialog = function(editor, element) {

	this.editor  = editor;
	this.element = element;
	this.element.setAttribute("id", "editor-open");

	this._documents = [];

	this.Init();

};

Dialog.prototype = {

	Init: function() {

		let create_file    = this.element.querySelector("input[name=\"create-file\"]");
		let open_file      = this.element.querySelector("select[name=\"open-file\"]");
		let cancel_button  = this.element.querySelector("button[data-action=\"cancel\"]")
		let confirm_button = this.element.querySelector("button[data-action=\"confirm\"]")

		if (create_file !== null && open_file !== null) {

			create_file.addEventListener("change", () => {

				if (confirm_button !== null) {

					if (is_valid_filepath(create_file.value)) {
						confirm_button.innerHTML = "Create";
					} else {
						confirm_button.innerHTML = "Open";
					}

				}

			});

			open_file.addEventListener("change", () => {

				if (confirm_button !== null) {

					if (is_valid_filepath(create_file.value)) {
						confirm_button.innerHTML = "Create";
					} else {
						confirm_button.innerHTML = "Open";
					}

				}

			});

		}

		if (cancel_button !== null) {

			cancel_button.addEventListener("click", () => {
				this.Hide();
			});

		}

		if (confirm_button !== null) {

			confirm_button.addEventListener("click", () => {

				let file_path = "";

				if (is_valid_filepath(create_file.value)) {
					file_path = create_file.value;
				} else if (open_file.value !== "") {
					file_path = open_file.value;
				}

				if (file_path !== "") {
					this.editor.Open(file_path);
				}

				this.Hide();

			});

		}

		this.element.addEventListener("click", (event) => {

			if (event.target === this.element) {
				this.Hide();
			}

		}, true);

	},

	Show: function() {

		this.editor.Config((config) => {

			this._documents = config.documents;

			let create_file = this.element.querySelector("input[name=\"create-file\"]");
			if (create_file !== null) {
				create_file.value = "";
			}

			let open_file = this.element.querySelector("select[name=\"open-file\"]");
			if (open_file !== null) {

				open_file.innerHTML = Object.keys(this._documents).sort((a, b) => {
					return a < b;
				}).map((file_path) => {

					let date = this._documents[file_path] || "0000-00-00";
					if (date !== "0000-00-00") {
						return "<option value=\"" + file_path + "\">[x] " + file_path + "</option>";
					} else {
						return "<option value=\"" + file_path + "\">[ ] " + file_path + "</option>";
					}

				});

			}

			this.element.setAttribute("open", "");

		});

	},

	Hide: function() {
		this.element.removeAttribute("open");
	}

};
