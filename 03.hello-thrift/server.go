package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"hello/gen-go/hello"
)

func main() {
	handler := &HelloServiceImpl{}
	processor := hello.NewHelloWorldServiceProcessor(handler)

	//创建服务器传输对象
	transportFactory := thrift.NewTTransportFactory()
	factoryDefault := thrift.NewTBinaryProtocolFactoryDefault()
	socket, err := thrift.NewTServerSocket(":9090")
	if err != nil {
		panic(err)
	}

	server4 := thrift.NewTSimpleServer4(processor, socket, transportFactory, factoryDefault)

	fmt.Printf("Starting the server...")

	server4.Serve()

}
