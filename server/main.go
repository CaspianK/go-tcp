package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fmt.Println("Server started on port 8080")
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		fmt.Println("New connection")

		go func() {
			for {
				buffer := make([]byte, 1024)
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println("Connection closed")
					return
				}

				messg := buffer[:n]

				_, err = conn.Write([]byte("Hello, " + string(messg)))
				if err != nil {
					log.Println(err)
					os.Exit(1)
				}
			}
		}()
	}
}
