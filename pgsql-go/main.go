package main

import (
	"log"

	"github.com/ddefrancesco/handson-go/pgsql-go/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		log.Println(err)
	}
	log.Println("Database tables are successfully initialized.")
}
