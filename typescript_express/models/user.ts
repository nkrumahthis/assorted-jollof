import db from "../database";

export type User = {
	id: number | undefined;
	name: string;
	email: string;
    password: string
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
		"INSERT INTO users (name, email, password) VALUES (:name, :email, password)"
	);

	const info = stmt.run(user);
	const data = get(Number(info.lastInsertRowid));

	return data;
}

export function update(id: number, user: User): User | undefined {
	const { name, email, password } = user;
	const sets: string[] = [];
	if (name) sets.push(`name = '${name}'`);
	if (email) sets.push(`email = '${email}'`);
	if (password) sets.push(`password = '${password}'`);
	const setClause = sets.join(", ");
	if (!setClause) return;

	const sql = `UPDATE users SET ${setClause} WHERE id = ${id}`;
	console.log(sql)
	const info = db.prepare(sql).run();
	const data = get(id);
	return data;
}

export function remove(id: number): boolean {
	const sql = `DELETE FROM users WHERE id = ${id}`;
	const info = db.prepare(sql).run()
	console.log(info)
	return info.changes > 0
}
