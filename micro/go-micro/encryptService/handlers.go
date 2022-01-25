package main

import (
	"context"

	proto "github.com/ddefrancesco/handson-go/micro/go-micro/encryptService/proto"
)

type Encrypter struct{}

func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, resp *proto.Response) error {
	resp.Result = EncryptString(req.Key, req.Message)
	return nil
}

func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, resp *proto.Response) error {
	resp.Result = DecryptString(req.Key, req.Message)
	return nil
}
