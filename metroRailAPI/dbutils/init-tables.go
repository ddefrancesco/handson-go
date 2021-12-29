package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(driver *sql.DB) {
	statement, driverError := driver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}
	//Create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!")
	}
	statement, _ = driver.Prepare(station)
	statement.Exec()
	statement, _ = driver.Prepare(schedule)
	statement.Exec()
	log.Println("All tables created/initialized successfully!")
}
