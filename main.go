package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()
	listener, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatal("Erro server on start")
	}

	defer listener.Close()
	log.Printf("Server Started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Erro ao acessar conexao")
			continue
		}

		go s.newClient(conn)
	}
}
