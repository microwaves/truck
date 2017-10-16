package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

var (
	target string
	port   int
)

func init() {
	flag.StringVar(&target, "target", "", "<host>:<port>")
	flag.IntVar(&port, "port", 4444, "port")
	flag.Parse()
}

func main() {
	signals := make(chan os.Signal, 1)
	stop := make(chan bool)
	signal.Notify(signals, os.Interrupt)

	go func() {
		for _ = range signals {
			fmt.Println("\nCtrl-C received, stopping...")
			stop <- true
		}
	}()

	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Could not listen on %d: %v", port, err)
	}
	fmt.Printf("Listening on %d\n", port)

	client, err := incoming.Accept()
	if err != nil {
		log.Fatal("Could not accept connection", err)
	}
	defer client.Close()
	fmt.Printf("Client %v successfully connected!\n", client.RemoteAddr())

	targetConn, err := net.Dial("udp", target)
	if err != nil {
		log.Fatal("Could not connect to target", err)
	}
	defer targetConn.Close()
	fmt.Printf("Connection to %v established!\n", targetConn.RemoteAddr())

	go func() { io.Copy(targetConn, client) }()
	go func() { io.Copy(client, targetConn) }()

	<-stop
}
