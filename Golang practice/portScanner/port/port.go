package port

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type ScanningPorts struct {
	port    string
	success string
}

func ScanPort(protocol string, endpoint string, port int, printOnlyOpen bool) ScanningPorts {
	response := ScanningPorts{port: protocol + " / " + strconv.Itoa(port)}
	request_url := endpoint + ":" + strconv.Itoa(port)
	connection, error := net.DialTimeout(protocol, request_url, 60*time.Second)
	if !printOnlyOpen {
		fmt.Println("-------------------------------------------------")
	}
	if error != nil {
		response.success = "Close"
		if !printOnlyOpen {
			fmt.Printf("The %v is %v\n", response.port, response.success)
		}
		return response
	} else {
		defer connection.Close()
		response.success = "Open"
		fmt.Printf("The %v is %v\n", response.port, response.success)
		return response
	}
}

func GetAllPorts(serverIp string, endPort int, printOnlyOpen bool) []ScanningPorts {
	var AllResponse []ScanningPorts
	for iter := 1; iter <= endPort; iter++ {
		AllResponse = append(AllResponse, ScanPort("tcp", serverIp, iter, printOnlyOpen))
		AllResponse = append(AllResponse, ScanPort("udp", serverIp, iter, printOnlyOpen))
	}
	return AllResponse
}

func GetAPort(endPoint string, protocol string, port int) {
	fmt.Printf("The port: %v / %v is %v", protocol, port, ScanPort(protocol, endPoint, port, false))
}
