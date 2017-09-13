package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)
var addr string = "127.0.0.1:9501"
var disconnect = make(chan bool)
func main() {
	if len(os.Args) != 2 {
		fmt.Println("参数错误")
		os.Exit(1)
	}
	name := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err)

	

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	defer conn.Close()

	fmt.Println("connect successfully")

	go receiveMessage(conn)

	go func () {
		for {
			var msg string
			fmt.Scanln(&msg)

			if ((msg == "quit") || (msg == "exit")) {
				break
			}
			conn.Write([]byte(name + ":" + msg + "\n"))
		}	
	}()

	<-disconnect
	os.Exit(0)
}

func receiveMessage(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')

		fmt.Printf(string(msg))
		if err != nil {
			break
		}
	}
	disconnect <- true
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}