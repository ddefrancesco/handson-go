package main

import (
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("encrypter"),
	)

	//Init parsa gli argomenti a riga di comando
	service.Init()

	//Registra l'handler
	//proto.RegisterEncrypterServer()
}
