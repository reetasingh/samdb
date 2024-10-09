package server

import (
	"fmt"
	"net"
	"samdb/pkg/core"
)

//var wg sync.WaitGroup

func talk(conn net.Conn) {
	for {
		cmd, err := core.ReadCmd(conn)
		if err != nil {
			conn.Close()
			//wg.Done()
			return
		}
		fmt.Println("Received cmd: ", cmd.Cmd, cmd.Args)
		result, err := core.ProcessCmd(cmd)
		if err != nil {
			_, err = conn.Write(core.EncodeError(err))
			if err != nil {
				conn.Close()
				//wg.Done()
				return
			}
		}
		_, err = conn.Write(core.EncodeString(result))
		if err != nil {
			conn.Close()
			//wg.Done()
			return
		}
	}
}

func AsyncTCPConnect(port int) error {
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
		fmt.Printf("client connected: %v\n", conn.RemoteAddr())
		fmt.Printf("concurrent connections: %d", connections)
		go talk(conn)
	}
}
