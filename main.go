package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: alarm [port]")
		return
	}

	port, e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Println("invalid port number")
		return
	}

	listenPort := fmt.Sprintf(":%d", port)
	fmt.Println("listen port", port)

	e = StartHTTP(listenPort)
	if e != nil {
		fmt.Println(e)
	}

}
