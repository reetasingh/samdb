package server

import (
	"fmt"
	"syscall"
)

func AsyncEpollTCPConnect(port int) error {
	// create server Socket
	serverFd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return err
	}
	// bind server Socket to Port and ipv4
	ipv4ByteArray := [4]byte{
		byte(127),
		byte(0),
		byte(0),
		byte(1),
	}
	sa := syscall.SockaddrInet4{Port: port, Addr: ipv4ByteArray}
	err = syscall.Bind(serverFd, &sa)
	if err != nil {
		return err
	}
	activeAllowedConnections := 10000

	// start listening
	err = syscall.Listen(serverFd, activeAllowedConnections)
	if err != nil {
		return err
	}

	fmt.Println("server started listening on: ", port)

	// Creates a new kernel event queue and returns a descriptor.
	kqFD, err := syscall.Kqueue()
	if err != nil {
		return err
	}

	//defer kqFD.close()

	changes := []syscall.Kevent_t{}
	// start monitoring serverFD first
	change := syscall.Kevent_t{}
	change.Ident = uint64(serverFd)
	change.Filter = syscall.EVFILT_READ
	change.Flags = syscall.EV_ADD | syscall.EV_ENABLE
	change.Fflags = 0
	change.Data = 0
	change.Udata = nil
	changes = append(changes, change)

	events := make([]syscall.Kevent_t, 1)

	for {
		// Used to register events with the queue, then wait for and return any pending
		// events to the user. In contrast to epoll, kqueue uses the same function to register
		// and wait for events, and multiple event sources may be registered and modified using a single call.
		// The changelist array can be used to pass modifications (changing the type of events to wait for, register new event sources, etc.) to the event queue,
		// which are applied before waiting for events begins. nevents is the size of the user supplied eventlist array that is used to receive events from the event queue.
		n, err := syscall.Kevent(kqFD, changes, events, nil)
		if err != nil {
			return err
		}
		fmt.Println("outer for loop")
		for i := 0; i < n; i++ {

			if events[i].Ident != uint64(serverFd) {
				// existing client connection
				// ReadAndEval(events[i].Udata)
				fmt.Println("existing client connection")
			} else {
				// new client connection
				// modiy change list to start monitoring this client
				change := syscall.Kevent_t{}
				change.Ident = events[i].Ident
				change.Filter = syscall.EVFILT_READ
				change.Flags = syscall.EV_ADD | syscall.EV_ENABLE
				change.Fflags = 0
				change.Data = 0
				change.Udata = nil
				changes = append(changes, change)
				fmt.Println("new client connection")
			}
		}
	}

}
