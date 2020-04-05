package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	var err error
	if l := len(os.Args) - 1; l != 2 {
		fmt.Printf("Received %d arguments,  client takes exactly 2 arguments: IP and Port\n", l)
		panic(-1)
	}
	ip := os.Args[1]
	port := os.Args[2]
	ipAddr := ip + ":" + port
	for {
		fmt.Println("Sending msg")
		time.Sleep(1 * time.Second)
		_, err = net.Dial("tcp", ipAddr)
		if err != nil {
			fmt.Printf("ERROR!!: %s\n", err)
		}
	}
}
