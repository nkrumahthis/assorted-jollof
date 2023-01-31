import express, { Request, Response } from "express";
import {
	ERROR_ORDER_NOT_FOUND,
	Order,
	all,
	create,
	get,
	remove,
	update,
} from "./order";

const router = express.Router();

// create
router.post("/", (req: Request, res: Response) => {
	let request = req.body;
	let response = create(request);
	res.status(201).json(response);
});

// read all
router.get("/", (req: Request, res: Response) => {
	let response = all();
	res.status(200).json(response);
});

// read one
router.get("/:id", (req: Request, res: Response) => {
	let id: number = Number(req.params.id);
	try {
		let response = get(id);
		if (response) return res.status(200).json(response);
	} catch (error) {
		if (error == ERROR_ORDER_NOT_FOUND) return res.status(404).json();
		else return res.status(500).json(error);
	}
});

// update
router.put("/:id", (req: Request, res: Response) => {
	let id: number = Number(req.params.id);
	let request: Order = req.body;
	let response = update(id, request);
	res.json(response);
});

// delete
router.delete("/:id", (req: Request, res: Response) => {
	let id: number = Number(req.params.id);
	try {
		let response = remove(id);
		if (response) res.status(204).json({});
	} catch (error) {
		if (error == ERROR_ORDER_NOT_FOUND) return res.status(404).json();
		else return res.status(500).json(error);
	}
});

export default router;
