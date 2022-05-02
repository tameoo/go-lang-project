package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	litenner, errListen := net.Listen("tcp", ":3000")

	if errListen != nil {
		fmt.Println(errListen)
		return
	}

	connection, errAccept := litenner.Accept()

	if errAccept != nil {
		fmt.Println(errAccept)
		return
	}

	for {
		data, errReader := bufio.NewReader(connection).ReadString('\n')

		if errReader != nil {
			fmt.Println(errReader)
			return
		}

		response := "Hello, " + string(data)
		connection.Write([]byte(response + "\n"))
	}
}
