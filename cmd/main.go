package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tada3/triton/handler"
)

const (
	port = 10600
)

func main() {
	log.Println("You won't get lost if you take the road to the left and bear to the left at every crossroad...")

	myport := port
	if flag.NArg() > 0 {
		var err error
		myport, err = strconv.Atoi(flag.Args()[0])
		if err != nil {
			fmt.Printf("Invalid port number: %s\n", flag.Args()[0])
			myport = port
		}
	}

	fmt.Printf("Listen on port %d\n", myport)

	http.HandleFunc("/cek", handler.Dispatch)

	addr := fmt.Sprintf(":%d", myport)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
