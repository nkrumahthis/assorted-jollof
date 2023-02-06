import db from "../database";

export type Customer = {
	id: number | undefined;
	name: string;
	phone: string;
	token: string
};

export const ERROR_CUSTOMER_NOT_FOUND = new Error("customer not found");
export const ERROR_NOT_IMPLEMENTED = new Error("not yet implemented");

export function all(): Customer[] {
	let data: Customer[] = db.prepare("SELECT * FROM customers").all();
	return data;
}

export function get(id: number): Customer {
	let data: Customer = db.prepare(`SELECT * FROM customers WHERE id=?`).get(id);
	if (!data) throw ERROR_CUSTOMER_NOT_FOUND;
	return data;
}

export function create(customer: Customer): Customer {
	// The following are equivalent.
	const stmt = db.prepare(
		"INSERT INTO customers (name, phone, token) VALUES (:name, :phone, :token)"
	);

	const info = stmt.run(customer);
	const data = get(Number(info.lastInsertRowid));

	return data;
}

export function update(id: number, customer: Customer): Customer | undefined {
	const { name, phone, token } = customer;
	const sets: string[] = [];
	if (name) sets.push(`name = '${name}'`);
	if (phone) sets.push(`phone = '${phone}'`);
	if (token) sets.push(`token = '${token}'`);
	const setClause = sets.join(", ");
	if (!setClause) return;

	const sql = `UPDATE customers SET ${setClause} WHERE id = ${id}`;
	console.log(sql)
	const info = db.prepare(sql).run();
	const data = get(id);
	return data;
}

export function remove(id: number): boolean {
	const sql = `DELETE FROM customers WHERE id = ${id}`;
	const info = db.prepare(sql).run()
	console.log(info)
	return info.changes > 0
}
