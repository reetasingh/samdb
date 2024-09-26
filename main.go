package main

import (
	"fmt"
	"net"
	"os"
	"samdb/pkg/core"
)

func connect(port int) error {
	laddr := net.TCPAddr{Port: port}
	listener, err := net.ListenTCP("tcp", &laddr)
	if err != nil {
		return err
	}
	connections := 0
	for {

		conn, err := listener.Accept()
		if err != nil {
			return nil
		}
		connections = connections + 1
		fmt.Printf("client connected: %v", conn.RemoteAddr())
		fmt.Printf("concurrent connections: %d", connections)

		for {
			cmd, err := core.ReadCmd(conn)
			if err != nil {
				conn.Close()
				break
			}
			fmt.Println(cmd.Cmd, cmd.Args)

			_, err = conn.Write([]byte("hello client"))
			if err != nil {
				conn.Close()
			}
		}
	}
}

func main() {
	err := connect(7379)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
