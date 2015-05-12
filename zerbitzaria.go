package main

import (
	"fmt"
	"net"
)

func zerbitzaria() *net.TCPListener {

	fmt.Println("Zerbitzaria abiarazten...\n")
	//zerbitzariaren konfigurazioa
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	erroreaAztertu(err)

	//zerbitzaria entzuten jarri
	ln, err := net.ListenTCP("tcp", tcpAddr)
	erroreaAztertu(err)

	return ln
}
