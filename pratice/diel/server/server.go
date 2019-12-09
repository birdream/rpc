package main

import (
	"fmt"
	"net"
	"time"
)

func handler(conn net.Conn, num int) {
	// defer conn.Close()

	// var (
	// 	message    string
	// 	newMessage string
	// )
	fmt.Println("<<<<<<<<<>>>>>>>>>")

	// message, _ = bufio.NewReader(conn).ReadString('\n')

	// fmt.Printf("Message Received: %s \n", string(message))

	// newMessage = strings.ToUpper(message)

	// conn.SetWriteDeadline(10 * time.Time)

	// conn.Write([]byte(newMessage + "\n"))
	message := fmt.Sprintf("%s %d \n", "anthor msg for", num)

	fmt.Println(message)

	conn.Write([]byte(message))
	time.Sleep(2 * time.Second)
}

func main() {
	var (
		ln   net.Listener
		err  error
		conn net.Conn
	)

	if ln, err = net.Listen("tcp", ":8080"); err != nil {
		fmt.Println("===================")
		fmt.Println(err)

		return
	}

	fmt.Println("tcp server is running..")

	if conn, err = ln.Accept(); err != nil {
		fmt.Println("++++++++++++++++++++")
		fmt.Println(err)

		return
	}

	var i int
	for {
		i++
		fmt.Printf("--going %d", i)
		handler(conn, i)
	}

}
