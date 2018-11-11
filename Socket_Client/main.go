package main

import (
	"flag"
	"log"
	"net"
	"os"
	"time"
)

var (
	ipAddress string
)

func main() {
	flag.StringVar(&ipAddress, "IPAddress", "1234", "IP Address for socket dial")
	flag.Parse()

	tcpAddress, err := net.ResolveTCPAddr("tcp", ipAddress)
	if err != nil {
		log.Fatal("IP Address resolve fail!")
		os.Exit(1)
	}
	log.Println("Ready to dial " + ipAddress)
	conn, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		log.Fatal("Dial " + ipAddress + " fail")
		os.Exit(1)
	}

	sender(conn)

}

func sender(conn net.Conn) {
	conn.Write([]byte("Hello Wold"))
	log.Fatal("Send at:" + time.Now().String())
}
