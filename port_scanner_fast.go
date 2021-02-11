package main

import (
	"fmt"
	"net"
)

func main() {
	// TCP ports range from 1 to 65535
	// For test scan 1-1024
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)

	}
}
