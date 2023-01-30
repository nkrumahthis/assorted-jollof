import db from "./database";

export type Order = {
	id: number | undefined;
	packs: number;
	user_id: number;
	location: string;
	status: string;
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
	const data = get(Number(info.lastInsertRowid))

	return data;
}

export function update(id: number, order: Order): Order {
	throw ERROR_NOT_IMPLEMENTED;
}

export function remove(id: number): boolean {
	throw ERROR_NOT_IMPLEMENTED;
}
