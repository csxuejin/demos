package main

import (
	"flag"
	"log"
	"net"
)

var(
	ip string
	port int
)

func init(){
	flag.StringVar(&ip, "ip", "127.0.0.1", "udp server endpoint")
	flag.IntVar(&port, "port", 9981, "udp server port")
}

func main(){
	addr := &net.UDPAddr{ IP: net.ParseIP(ip),Port: port}
	listener , err := net.ListenUDP("udp", addr)
	if err != nil{
		panic(err)
		return
	}

	log.Printf("udp server: %v\n", listener.LocalAddr())

	data := make([]byte, 1024)
	for{
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil{
			log.Printf("litener.ReadFromUDP(): %v\n", err)
			continue
		}

		log.Printf("<%s> %s\n", remoteAddr, data[:n])

		go serve(listener, remoteAddr, data[:n])
	}
}

func serve(listener *net.UDPConn, remoteAddr *net.UDPAddr, data []byte) {
	_, err := listener.WriteToUDP(data, remoteAddr)
	if err != nil{
		log.Printf("listener.WriteToUDP(): %v\n", err)
	}
}
