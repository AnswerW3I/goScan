package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Print("test")
}

func tcpScan(ip string) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Print("Connect Sucessful")
	conn.Close()
}
