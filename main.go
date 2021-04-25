package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ip, err := OutboundIP()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ip)
}

// OutboundIP returns string form of address (for example, "192.0.2.1", "[2001:db8::1]")
func OutboundIP() (ip string, err error) {
	conn, err := net.Dial("ip:icmp", "google.com")
	if err != nil {
		conn, err = net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			return
		}
	}
	ip = conn.LocalAddr().String()
	conn.Close()
	a := strings.Split(ip, ":")
	ip = a[0]
	return
}
