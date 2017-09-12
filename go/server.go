package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main() {
	addr := "127.0.0.1:9501"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		reader := bufio.NewReader(conn)

		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				break;
			}
			fmt.Println(string(msg))
			conn.Write([]byte("hello client\n"))
		}
		conn.Close()
	}


}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}