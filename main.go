package main

import (
	"fmt"
	"net"
	"os"
)

func connect(port int) error {
	laddr := net.TCPAddr{Port: port}
	listener, err := net.ListenTCP("tcp", &laddr)
	if err != nil {
		return err
	}

	for {
		fmt.Println("listening")
		conn, err := listener.Accept()
		if err != nil {
			return nil
		}

		fmt.Println(conn.LocalAddr())
	}
}

func main() {
	err := connect(7379)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
