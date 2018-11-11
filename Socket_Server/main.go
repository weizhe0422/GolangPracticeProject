package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var (
	ipAddress string
)

func main() {
	flag.StringVar(&ipAddress, "IPAddress", "1234", "IP Address for socket listen")
	flag.Parse()

	listener, err := net.Listen("tcp", ipAddress)
	if err != nil {
		log.Fatal("Listen " + ipAddress + " fail")
	}

	log.Println("Listen at " + ipAddress)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleFunc(conn)
	}

}

func handleFunc(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}
	daytime := time.Now().String()

	//conn.Write([]byte("Recieve at: " + daytime + ", message:" + string(buffer[:n])))
	log.Println("Recieve at: " + daytime + ", message:" + string(buffer[:n]))
}
