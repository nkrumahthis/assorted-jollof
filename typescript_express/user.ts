import db from "./database";

export type User = {
	id: number | undefined;
	name: string;
	phone: string;
};

export const ERROR_USER_NOT_FOUND = new Error("user not found");
export const ERROR_NOT_IMPLEMENTED = new Error("not yet implemented");

export function all(): User[] {
	let data: User[] = db.prepare("SELECT * FROM users").all();
	return data;
}

export function get(id: number): User {
	let data: User = db.prepare(`SELECT * FROM users WHERE id=?`).get(id);
	if (!data) throw ERROR_USER_NOT_FOUND;
	return data;
}

export function create(user: User): User {
	// The following are equivalent.
	const stmt = db.prepare(
		"INSERT INTO users (name, phone) VALUES (:name, :phone)"
	);

	const info = stmt.run(user);
	const data = get(Number(info.lastInsertRowid))

	return data;
}

export function update(id: number, user: User): User {
	throw ERROR_NOT_IMPLEMENTED;
}

export function remove(id: number): boolean {
	throw ERROR_NOT_IMPLEMENTED;
}
