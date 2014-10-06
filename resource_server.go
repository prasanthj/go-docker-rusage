package main

import (
	"net"
	"log"
	"syscall"
	"bytes"
	"encoding/gob"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8787"
	CONNECTION_TYPE = "tcp"
)

func main() {
	
	// listen for incoming connections
	l, err := net.Listen(CONNECTION_TYPE, ":" + CONNECTION_PORT)
	if err != nil {
		log.Fatalf("Error listening to connections: %s\n", err.Error())
	}

	// close listener when the application exits
	defer l.Close()

	log.Println("Listening on " + CONNECTION_HOST + ":" + CONNECTION_PORT)

	// keeps the serivce running forever
	for {

		// accept incoming connections
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Error accepting connections: %s", err.Error())
		}

		log.Printf("Received message from %s\n", conn.RemoteAddr())

		// the resource usage struct has to serialized before sending over
		// the wire. For serialization, we must register the concrete type
		// during encoding and decoding. On each end, this tells the
		// engine which concrete type is being sent that implements the interface.
		var buf bytes.Buffer
		gob.Register(syscall.Rusage{})
		encoder := gob.NewEncoder(&buf)

		// handle request in a go routine
		go handleRequest(conn, encoder, &buf)
	}
}

func handleRequest(conn net.Conn, encoder *gob.Encoder, buf* bytes.Buffer) {

	// create buffer and read data from incoming connection. In this resource
	// usage example the client doesn't really send anything useful for the server.
	rbuf := make([]byte, 1024)
	if _, err := conn.Read(rbuf); err != nil {
		panic(err)
	}

	// get the resource usage of self (current process)
	var rusage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &rusage); err != nil {
		panic(err)
	}

	// serialize the resource usage struct and write it to connection
	err := encoder.Encode(&rusage)
	if err != nil {
		panic(err)
	}
	conn.Write(buf.Bytes())

	// close the connection
	conn.Close()
}