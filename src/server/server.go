package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting server")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		msg, err := ln.Accept()
		if err!= nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
	}
}