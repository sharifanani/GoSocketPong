package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var err error
	for {
		fmt.Println("Sending msg")
		time.Sleep(1 * time.Second)
		_, err = net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}
	}
}