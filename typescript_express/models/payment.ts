import db from "../database";

export type Payment = {
	id: number;
	amount: number;
	orderId: number;
	timestamp: string;
	userId: number;
};

export const ERROR_PAYMENT_NOT_FOUND = new Error("payment not found");
export const ERROR_NOT_IMPLEMENTED = new Error("not yet implemented");

export function all(): Payment[] {
	let data: Payment[] = db.prepare("SELECT * FROM payments").all();
	return data;
}

export function get(id: number): Payment {
	let data: Payment = db.prepare(`SELECT * FROM payments WHERE id=?`).get(id);
	if (!data) throw ERROR_PAYMENT_NOT_FOUND;
	return data;
}

export function create(payment: Payment): Payment {
	// The following are equivalent.
	const stmt = db.prepare(
		"INSERT INTO payments (amount, orderId, userId) VALUES (:amount, :orderId, :userId)"
	);

	const info = stmt.run(payment);
	const data = get(Number(info.lastInsertRowid));

	return data;
}

export function update(id: number, payment: Payment): Payment | undefined {
	const { amount, orderId, userId } = payment;
	const sets: string[] = [];
	if (amount) sets.push(`amount = '${amount}'`);
	if (orderId) sets.push(`orderId = '${orderId}'`);
	if (userId) sets.push(`userId = '${userId}'`);
	const setClause = sets.join(", ");
	if (!setClause) return;

	const sql = `UPDATE payments SET ${setClause} WHERE id = ${id}`;
	console.log(sql);
	const info = db.prepare(sql).run();
	const data = get(id);
	return data;
}

export function remove(id: number): boolean {
	const sql = `DELETE FROM payments WHERE id = ${id}`;
	const info = db.prepare(sql).run();
	console.log(info);
	return info.changes > 0;
}
