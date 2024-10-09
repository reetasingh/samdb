package main

import (
	"fmt"
	"os"
	"samdb/pkg/server"
)

func main() {
	err := server.AsyncKqueueTCPConnect(7380)
	//err := server.SyncTCPConnect(7380)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
