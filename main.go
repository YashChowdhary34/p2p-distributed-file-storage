package main

import (
	"log"

	"github.com/YashChowdhary34/p2p-distributed-file-storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err!=nil{
		log.Fatal(err)
	}

	select{}
}