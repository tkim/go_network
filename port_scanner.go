package main

import (
	"fmt"
	"net"
)

func main() {
	// TCP ports range from 1 to 65535
	// For test scan 1-1024
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org: %d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
