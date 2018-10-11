package main

import (
	"fmt"
	"github.com/KevinZu/fins"
)

func main() {
	var error_val int32
	error_val = 0
	sys := new(fins.FinsSysTp)
	cliInfo, err := sys.FinslibTcpConnect("192.168.1.1", 9600, 0, 10, 0, 0, 1, 0, &error_val, 6)
	if err != nil || error_val != 0 {
		fmt.Printf("FinslibTcpConnect error ! error_val:%d\n", error_val)
	}

	sys.CliGroup.DelClient(cliInfo)

	fmt.Printf("FinslibTcpConnect OK ! \n")
}
