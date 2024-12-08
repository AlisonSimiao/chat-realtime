package database

import (
	"database/sql"
	"log"
)

// DB is a global variable for the SQLite database connection
var DB *sql.DB

// initDB initializes the SQLite database and creates the todos table if it doesn't exist
func initDB() {
 var err error
 DB, err = sql.Open("sqlite3", "./app.db") // Open a connection to the SQLite database file named app.db
 if err != nil {
  log.Fatal(err) // Log an error and stop the program if the database can't be opened
 }

 sqlStmt := `
 CREATE TABLE IF NOT EXISTS rooms (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  color integer NOT NULL,
  name TEXT
 );
 
 CREATE TABLE IF NOT EXISTS colors (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  hex TEXT
 );
 `

 _, err = DB.Exec(sqlStmt)
 if err != nil {
  log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt) // Log an error if table creation fails
 }
}