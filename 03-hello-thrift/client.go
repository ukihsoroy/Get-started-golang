package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"hello/gen-go/hello"
)

func main() {
	socket, err := thrift.NewTSocket(":9090")
	if err != nil {
		panic(err)
	}

	//创建客户端链接
	factoryDefault := thrift.NewTBinaryProtocolFactoryDefault()
	client := hello.NewHelloWorldServiceClientFactory(socket, factoryDefault)

	//打开链接
	if err := socket.Open(); err != nil {
		panic(err)
	}
	defer socket.Close()

	r, err := client.Say(context.Background(), "K.O")

	if err != nil {
		panic(nil)
	}
	fmt.Println(r)
}
