package app

import (
	"database/sql"
	"log"

	"gopkg.in/gorp.v1"
)

// Database struct
// type Database gorp.DbMap
//
// var database *Database
var db *gorp.DbMap

func checkErr(err error, msg string) {

	if err != nil {
		log.Fatalln(msg, err)
	}
}

// InitDB func
func InitDB() {

	dbConn, err := sql.Open("mysql", "wyma:1234@tcp(192.168.99.100:3306)/myapp")
	checkErr(err, "sql.Open failed")

	dialect := gorp.MySQLDialect{"InnoDB", "UTF8"}
	db = &gorp.DbMap{Db: dbConn, Dialect: dialect}

	return
}

// AddTables func
// add a table for Database
func AddTables(i interface{}, name string) {

	db.AddTableWithName(i, name).SetKeys(true, "ID")
	err := db.CreateTablesIfNotExists()

	checkErr(err, "Create tables failed")

	return
}

// GetDB func
func GetDB() *gorp.DbMap {
	return db
}
