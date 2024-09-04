package conn

import "net"

func connect(port int) {
	laddr := net.TCPAddr{Port: port}
	net.ListenTCP("localhost", &laddr)

}
