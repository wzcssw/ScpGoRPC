package main

import (
	"goRPC/lib"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		lib.Receive()
	} else {
		lib.Send(os.Args[1], os.Args[2])
	}
}
