package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

var ConnMap map[string]*net.TCPConn

func main() {
	addr := "127.0.0.1:9501"
	ConnMap = make(map[string]*net.TCPConn)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	listener, err := net.ListenTCP("tcp4", tcpAddr)
	checkError(err)

	defer listener.Close()


	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}

		ipAddr := conn.RemoteAddr().String()

		ConnMap[ipAddr] = conn

		go handleConn(conn)
	}


}

func handleConn(conn *net.TCPConn) {
	// ipAddr := conn.RemoteAddr().String()

	reader := bufio.NewReader(conn)

	defer conn.Close()
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(string(msg))

		for _,other := range ConnMap {
			other.Write([]byte(msg))
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}