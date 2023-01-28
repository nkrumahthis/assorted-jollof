export type User = {
	id: number | undefined;
	name: string;
	phone: string;
};

export const ERROR_USER_NOT_FOUND = new Error("user not found");

let data: User[] = [
	{ id: 0, name: "Dana", phone: "080323789010" },
	{ id: 1, name: "Eric", phone: "02033342312" },
	{ id: 2, name: "Frank", phone: "02033342312" },
	{ id: 3, name: "Gerald", phone: "02033342312" },
	{ id: 4, name: "Harry", phone: "02033342312" },
];

export function all(): User[] {
	return data;
}

export function get(id: number): User {
	let found = data.some((item) => item.id == id);
	if (!found) throw ERROR_USER_NOT_FOUND;

	return data[id];
}

export function create(user: User): User {
	user.id = data.length - 1;
	data.push(user);
	return data[data.length - 1];
}

export function update(id: number, user: User): User {
	let found = data.some((item) => item.id == id);
	if (!found) throw ERROR_USER_NOT_FOUND;

	data = data.filter((item) => item.id !== id);
	return data[id];
}

export function remove(id: number): boolean {
	let found = data.some((item) => item.id == id);
	if (!found) throw ERROR_USER_NOT_FOUND;

	data = data.filter((item) => item.id !== id);
	return found;
}
