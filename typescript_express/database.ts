import Database from "better-sqlite3";
import { readFileSync } from "fs";
import {resolve} from "path"

// https://sagot.dev/en/articles/sqlite-typescript/

console.log("setting up db")

// create tables if they don't exist
const CREATE_TABLES = readFileSync(resolve("..", "main.sql")).toString();
const SEED_DATA = readFileSync(resolve("..", "seed.sql")).toString();

const db = new Database("../main.db");
db.pragma('journal_mode = WAL');

db.exec(CREATE_TABLES);
db.exec(SEED_DATA);

export default db;
