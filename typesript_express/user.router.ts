import express from "express";
import { all, get } from "./user";
const router = express.Router();

// define the home page route
router.get("/", (req, res) => {
    let response = all()
	res.json(response);
});

router.post("/", (req, res) => {
	res.send("user home page");
});

router.put("/:id", (req, res) => {
    let id = req.params.id;
	res.send("user home page");
});


// define the about route
router.get("/:id", (req, res) => {
    let id = req.params.id;
    let response = get(id)
	res.json(response);
});

export default router;
