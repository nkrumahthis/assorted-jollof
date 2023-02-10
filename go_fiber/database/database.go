package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func PrepDB() {

	mainsql, err := os.ReadFile("../main.sql")
	if err != nil {
		panic(err)
	}

	seedsql, err := os.ReadFile("../seed.sql")
	if err != nil {
		panic(err)
	}

	CREATE_TABLES := string(mainsql)
	SEED_DATA := string(seedsql)

	db, err := GetDB()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(CREATE_TABLES)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(SEED_DATA)
	if err != nil {
		panic(err)
	}

	fmt.Println("database prepared")

}

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "../main.db")
	return db, err
}
