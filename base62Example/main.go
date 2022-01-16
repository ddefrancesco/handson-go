package main

import (
	"log"

	"github.com/ddefrancesco/handson-go/base62Example/base62"
)

func main() {
	x := 123456
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	normalNumber := base62.ToBase10(base62String)
	log.Println(normalNumber)
}
