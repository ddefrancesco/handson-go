package main

import (
	"log"

	"github.com/ddefrancesco/handson-go/micro/go-micro/encryptString/utils"
)

// AES keys should be of length 16, 24, 32
func main() {
	key := "111023043350789514532147"

	message := "Io sono un messaggio"
	log.Println("Original message: ", message)
	encryptedString := utils.EncryptString(key, message)
	log.Println("Encrypted message: ", encryptedString)
	decryptedString := utils.DecryptString(key, encryptedString)
	log.Println("Decrypted message: ", decryptedString)
}
