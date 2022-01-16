package helper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "git_user"
	password = "P4p3r1n0"
	dbname   = "mydb"
)

func InitDB() (*sql.DB, error) {
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Printf("Initializing DB %s", dbname)
	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	log.Printf("Open DB %s", dbname)
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS web_url(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
	if err != nil {
		return nil, err
	}
	log.Printf("Created DB statement: %s", stmt)
	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	log.Printf("Executed DB statement: %s", stmt)
	log.Println("Exiting InitDB func")
	return db, nil

}
