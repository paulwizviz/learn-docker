package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	nis, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, ni := range nis {
		addrs, err := ni.Addrs()
		if err != nil {
			break
		}
		for _, addr := range addrs {
			fmt.Printf("Name: %v Addr: %v\n", ni.Name, addr)
		}
	}
}
