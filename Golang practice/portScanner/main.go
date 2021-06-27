package main

import (
	"github.com/codertcet111/portScanner/port"
)

func main() {
	serverEndPoint := "127.0.0.1"
	endPort := 1024
	printOnlyOpen := true
	port.GetAllPorts(serverEndPoint, endPort, printOnlyOpen)

	port.GetAPort(serverEndPoint, "tcp", 5432)
}
