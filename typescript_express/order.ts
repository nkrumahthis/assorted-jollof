import db from "./database";

const PENDING = 'PENDING';
const FULFILLED = 'FULFILLED';
const CANCELLED = 'CANCELLED';

type OrderStatus = typeof PENDING | typeof FULFILLED | typeof CANCELLED;

export type Order = {
	id: number | undefined;
	packs: number;
	user_id: number;
	location: string;
	status: OrderStatus;
};

export const ERROR_ORDER_NOT_FOUND = new Error("order not found");
export const ERROR_NOT_IMPLEMENTED = new Error("not yet implemented");

export function all(): Order[] {
	let data: Order[] = db.prepare("SELECT * FROM orders").all();
	return data;
}

export function get(id: number): Order {
	let data: Order = db.prepare(`SELECT * FROM orders WHERE id=?`).get(id);
	if (!data) throw ERROR_ORDER_NOT_FOUND;
	return data;
}

export function create(order: Order): Order {
	// The following are equivalent.
	const stmt = db.prepare(
		"INSERT INTO orders (packs, user_id, location, status, phone) VALUES (:packs, :user_id, :location, :status, :phone)"
	);

	const info = stmt.run(order);
	const data = get(Number(info.lastInsertRowid));

	return data;
}

export function update(id: number, order: Order): Order | undefined {
	const { packs, user_id, location, status } = order;
	const sets: string[] = [];
	if (packs) sets.push(`packs = ${packs}`);
	if (user_id) sets.push(`user_id = ${user_id}`);
	if (location) sets.push(`location = '${location}'`);
	if (status) sets.push(`status = '${status}'`);

	const setClause = sets.join(", ");
	if (!setClause) return;

	const sql = `UPDATE orders SET ${setClause} WHERE id = ${id}`;

	const info = db.prepare(sql).run(order);
	const data = get(Number(info.lastInsertRowid));

	return data;
}

export function remove(id: number): boolean {
	const sql = `DELETE FROM orders WHERE id = ${id}`;
	const info = db.prepare(sql).run()
	console.log(info)
	return info.changes > 0
}
