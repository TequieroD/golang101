package main

import "database/sql"  
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "root@/godb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO test (name) VALUES(?)")
	
	stmtIns.Exec("yoyo")
	
	defer stmtIns.Close() 
}