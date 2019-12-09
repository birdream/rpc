package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	var (
		conn   net.Conn
		err    error
		res    string
		reader *bufio.Reader
		// isPrefix bool
		// byteRes []byte
		message string
	)

	if conn, err = net.Dial("tcp", "127.0.0.1:8080"); err != nil {
		fmt.Println("|||||||||||||")
		fmt.Println(err)
		return
	}

	for {
		reader = bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")

		res, _ = reader.ReadString('\n')
		fmt.Fprintf(conn, res+"\n")

		message, _ = bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message from server: %s", message)

		// time.Sleep(1000 * time.Second)
	}

	// fmt.Println("<<<<<<<<<<<<<")
	// fmt.Println(byteRes)
	// fmt.Println(string(byteRes))

	// if res, err = reader.ReadString('\n'); err != nil {
	// 	fmt.Println("<<<<<<<<<<<<<")
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println(res)
}
