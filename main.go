package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var MySQLWriter *sql.DB
var MySQLReader *sql.DB

func main() {
	openMySQLWriter()
	openMySQLReader()

	setUserEmail("b@example.com", "1")
	println(getUserEmail("1"))
}

// using MySQLWriter
func setUserEmail(email string, id string) {
	stmt, err := MySQLWriter.Prepare("update users set email=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(email, id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

// using MySQLReader
func getUserEmail(id string) string {
	var email string
	row := MySQLReader.QueryRow("SELECT email FROM users where id = " + id)
	e := row.Scan(&email)
	checkErr(e)
	return email
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
