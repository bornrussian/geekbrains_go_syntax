//
// Задача:
//
// Добавьте для time-сервера возможность его завершения при вводе команды ​ exit​ .
//

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	go listenConn()

	commandLineInterface()
}

func commandLineInterface() {
	var command string
	for {
		fmt.Scan(&command)
		switch command {
		case "exit":
			return
		default:
			fmt.Println("commands:")
			fmt.Println("  exit = terminate this server")
		}
	}
}

func listenConn() {
	fmt.Println("Listen on localhost:8123 ...")
	fmt.Println("You can: run 'telnet 127.0.0.1 8123'")
	fmt.Println("You can: terminate this server with 'exit' command")
	listener, err := net.Listen("tcp", "localhost:8123")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConn(conn)
	}
}


func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c,time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}