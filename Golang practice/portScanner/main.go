package main

import (
	"github.com/codertcet111/portScanner/port"
)

func main() {
	// endPort := 1024
	// printOnlyOpen := true
	// port.GetAllPorts("52.219.62.94", endPort, printOnlyOpen)

	port.GetAPort("52.219.62.94", "tcp", 5432)
}
