package main

import (
	"fmt"

	"github.com/hasael-web/Green_Goods_Market-BE/cmd"
)

func main() {
	go cmd.ServerGrpcGateway()
	cmd.ServerGrpc()
	fmt.Println("golang grpc run")
}
