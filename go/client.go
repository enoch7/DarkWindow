package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)
var addr string = "127.0.0.1:9501"

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

	_, err = conn.Write([]byte(name + ":Hello Server\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)


	
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}