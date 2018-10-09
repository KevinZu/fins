package main

import (
	"fmt"
	"github.com/KevinZu/fins"
)

func main() {
	var error_val int32
	sys := new(fins.FinsSysTp)
	err := sys.FinslibTcpConnect("192.168.1.1", 9600, 0, 10, 0, 0, 1, 0, &error_val, 6)
	if err != nil {
		fmt.Printf("FinslibTcpConnect error ! \n")
	}

	fmt.Printf("FinslibTcpConnect OK ! \n")
}
