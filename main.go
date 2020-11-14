package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sqlx.DB

	db, err := sqlx.Open("sqlite3", ":memory:")

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	createSchema(db)
}

// User ...
type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Lastname string `db:"lastname"`
}

func createSchema(db *sqlx.DB) {
	schema := `CREATE TABLE user (
			id       integer     PRIMARY KEY AUTOINCREMENT,
			name     varchar(56),
			lastname varchar(56)
		);`

	result, err := db.Exec(schema)
	if err != nil {
		fmt.Printf("\nError:\n\t%v\n", err)
		panic(err.Error())
	}
	fmt.Printf("\nResultado de creacion de la tabla user:\n \t%v\n", result)

	insertUser := `INSERT INTO user (name, lastname) VALUES (?, ?)`

	db.MustExec(insertUser, "Matias", "Gomez")

	rows, err := db.Queryx(`SELECT
														id,
														name,
														lastname
													FROM user`)
	if err != nil {
		fmt.Printf("\nError en INSERT:\n\t%v\n", err)
		panic(err.Error())
	}

	for rows.Next() {
		var user User

		err = rows.StructScan(&user)

		if err != nil {
			fmt.Printf("\nError en INSERT:\n\t%v\n", err)
		}
		fmt.Printf("\nResultado SELECT a tabla user:\n\t%v\n", user)
	}

}
