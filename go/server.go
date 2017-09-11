package main

import (
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

func main() {
	addr := "127.0.0.1:9501"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		result, err := ioutil.ReadAll(conn)
		if err != nil {
			conn.Close()
		}
		fmt.Println(string(result))
		conn.Write([]byte("hello client\r\n\r\n"))
		conn.Close()
	}


}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}