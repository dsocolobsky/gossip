package main

import (
	"bufio"
	"fmt"
	"net/textproto"
)

func main() {
	mcon := CreateConnection()
	sock, err := mcon.Connect()

	if err != nil {
		fmt.Printf("Error connecting to the server")
	}

	mcon.Start(sock)

	reader := bufio.NewReader(sock)
	tp := textproto.NewReader(reader)
	// The input system is really ugly and cumbersome
	// I use holder as fix, but should be rewritten
	input := ""
	holder := ""
	for input != "/quit" {
		line, err := tp.ReadLine()
		if err != nil {
			break // break loop on errors
		}
		fmt.Printf("%s\n", line)

		fmt.Scanf("%s", &input)
		if input != "" && input != holder {
			mcon.SendMessage(sock, input)
			holder = input
		}
	}
	
	return
}
