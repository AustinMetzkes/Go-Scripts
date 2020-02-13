/*
This is from Chapter 2 of Black Hat Go
Most of this (if not all) code has been copied from there
I may modify it slightly but credit for the code goes to the author
*/
package main

import (
	"fmt"
	"net"

)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//port is closed or filtered
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
