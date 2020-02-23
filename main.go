package main

import (
	"fmt"
	"github.com/babadee08/athena/port"
)

func main() {
	fmt.Println("Port scanner in GO")

	open := port.ScanPort("tcp", "localhost", 3306)

	fmt.Printf("Port open: %t\n", open)
}
