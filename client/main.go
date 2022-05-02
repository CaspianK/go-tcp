package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("– ")
		messg, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(messg + "\n"))
		if err != nil {
			log.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Connection closed")
			return
		}

		data := buffer[:n]
		fmt.Println(string(data))
	}
}
