export type User = {
	id: number;
	name: string;
	phone: string;
};

let data = [
	{ id: 0, name: "Dana", phone: "080323789010" },
	{ id: 1, name: "Eric", phone: "02033342312" },
	{ id: 2, name: "Frank", phone: "02033342312" },
	{ id: 3, name: "Gerald", phone: "02033342312" },
	{ id: 4, name: "Harry", phone: "02033342312" },
];

export function all(): User[] {
	return data;
}

export function get(id:number): User {
	return data[id];
}
