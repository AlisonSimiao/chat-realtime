package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB is a global variable for the SQLite database connection
var DB *sql.DB

// initDB initializes the SQLite database and creates the todos table if it doesn't exist
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db") // Open a connection to the SQLite database file named app.db
	if err != nil {
		log.Fatal(err) // Log an error and stop the program if the database can't be opened
	}

	sqlStmt := `
 CREATE TABLE IF NOT EXISTS colors (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE,
    hex TEXT UNIQUE
);

	CREATE TABLE IF NOT EXISTS rooms (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    id_color INTEGER NOT NULL,
    FOREIGN KEY (id_color) REFERENCES colors(id)
);

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		text_color INTEGER NOT NULL,
		background_color INTEGER NOT NULL,
		FOREIGN KEY (text_color) REFERENCES colors(id),
		FOREIGN KEY (background_color) REFERENCES colors(id)
	);

	INSERT INTO colors (name, hex) VALUES ('black', '#000000') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('white', '#ffffff') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('red', '#ff0000') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('green', '#00ff00') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('blue', '#0000ff') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('yellow', '#ffff00') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('magenta', '#ff00ff') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('cyan', '#00ffff') ON CONFLICT DO NOTHING;
    INSERT INTO colors (name, hex) Values ('gray', '#808080') ON CONFLICT DO NOTHING;
 `

	fmt.Println("criando db")
	result, err := DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %s: %s\n", err, result) // Log an error if table creation fails
	}
}
