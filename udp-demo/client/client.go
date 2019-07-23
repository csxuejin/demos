package main

import (
	"flag"
	"log"
	"net"
)

var (
	ip, port string
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "udp server endpoint")
	flag.StringVar(&port, "port", "9981", "udp server port")
}

func main() {
	service := ip + ":" + port
	remoteAddr, err := net.ResolveUDPAddr("udp", service)
	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	// write a message to server
	if _, err := conn.Write([]byte("Hello UDP server!")); err != nil {
		log.Println(err)
	}

	// receive message from server
	buffer := make([]byte, 1024)
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil{
		log.Println(err)
		return
	}

	log.Println("UDP Server : ", addr)
	log.Println("Received from UDP server : ", string(buffer[:n]))
}
