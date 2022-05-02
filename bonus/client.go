package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	connection, errDial := net.Dial("tcp", "127.0.0.1:3000")

	if errDial != nil {
		fmt.Println(errDial)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(connection, text+"\n")

		data, errReader := bufio.NewReader(connection).ReadString('\n')
		if errReader != nil {
			fmt.Println(errReader)
			return
		}

		fmt.Print(data)
	}
}
