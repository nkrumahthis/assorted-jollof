package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ac (x []byte, err error) []byte {
	if err != nil {
		panic(err)
	}

	return x
}

func PrepDB() {
	
	mainsql := ac(os.ReadFile("../main.sql"))
	
	seedsql := ac(os.ReadFile("../seed.sql"))
	
	CREATE_TABLES := string(mainsql)
	SEED_DATA := string(seedsql)
		
	db, err := sql.Open("sqlite3", "../main.db")
	if err != nil {
		panic(err)
	}
	
	db.Exec(CREATE_TABLES)
	db.Exec(SEED_DATA)

	fmt.Println("database prepared")

}

func GetDB() (*sql.DB, error){
	db, err := sql.Open("sqlite3", "../main.db")
	return db, err
}