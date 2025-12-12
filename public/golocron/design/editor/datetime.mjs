
const format = (str, length) => {

	if (str.length < length) {
		return new Array(length - str.length).fill('0').join('') + str;
	}

	return str;

};

export const datetime = () => {

	let name = '';
	let date = new Date();

	name += date.getFullYear();
	name += '-';
	name += format('' + (date.getUTCMonth() + 1), 2);
	name += '-';
	name += format('' + date.getUTCDate(), 2);

	return name;

};

