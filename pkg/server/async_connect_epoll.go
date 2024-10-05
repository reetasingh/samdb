package server

import (
	"fmt"
	"net"
	"samdb/pkg/core"

	"golang.org/x/sys/unix"
)

func AsyncEpollTCPConnect(port int) error {
	///laddr := net.TCPAddr{Port: port}
	listener, err := net.Listen("tcp", string(port))
	if err != nil {
		return err
	}
	defer listener.Close()

	file, err := listener.(*net.TCPListener).File()
	if err != nil {
		fmt.Println("Error getting file descriptor:", err)
	}

	// Create an epoll instance
	kq, err := unix.Kqueue()
	if err != nil {
		fmt.Errorf("Error creating epoll:%v", err)
	}
	defer unix.Close(kq)

	// Prepare the kqueue event
	event := unix.Kevent_t{
		Ident:  uint64(file.Fd()), // file descriptor to monitor
		Filter: unix.EVFILT_READ,
		Flags:  unix.EV_ADD | unix.EV_ENABLE,
		// Set additional fields as needed
	}

	// Add the event to kqueue
	if _, err := unix.Kevent(kq, []unix.Kevent_t{event}, nil, nil); err != nil {
		fmt.Errorf("Error adding event to kqueue:", err)
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

		for {
			cmd, err := core.ReadCmd(conn)
			if err != nil {
				conn.Close()
				break
			}
			fmt.Println("Received cmd: ", cmd.Cmd, cmd.Args)
			result, err := core.ProcessCmd(cmd)
			if err != nil {
				_, err = conn.Write(core.EncodeError(err))
				if err != nil {
					conn.Close()
				}
			}
			_, err = conn.Write(core.EncodeString(result))
			if err != nil {
				conn.Close()
			}
		}
	}
}
