package main

import (
	"net"
	"syscall"
	"bytes"
	"os"
	"io"
	"log"
	"encoding/gob"
	"encoding/json"
	"time"
)

const (
	SERVICE_TYPE = "tcp"
	SERVICE_HOST = "localhost"
	SERVICE_PORT = "8787"
	DIAL_TIMEOUT = time.Duration(5 * time.Second)
)

func main() {

	// if host not passed as 2nd argument use localhost
	var host = SERVICE_HOST
	if len(os.Args) == 2 {
		host = os.Args[1]
	}

	// connect to the desired host:port or timeout after 5 seconds
	log.Println("Dialing " + host + ":" + SERVICE_PORT)
	conn, err := net.DialTimeout(SERVICE_TYPE, host + ":" + SERVICE_PORT, DIAL_TIMEOUT)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to " + host + ":" + SERVICE_PORT)

	// close the connection when program exits
	defer conn.Close()

	// write line break to the request (blank request)
	if _, err := conn.Write([]byte("\n")); err != nil {
		panic(err)
	}
	log.Println("Blank request sent to server")

	// we expect syscall.Rusage struct from the server. To deserialize
	// the response from server we have to register the concrete type.
	gob.Register(syscall.Rusage{})

	// create buffer to hold the response and create the decoder
	var buf bytes.Buffer
	decoder := gob.NewDecoder(&buf)

	// copy the response from connection
	l, err := io.Copy(&buf, conn)
	if err != nil {
		panic(err)
	}
	log.Printf("Read %v bytes from resource service\n", l)

	// decode the response
	var result syscall.Rusage
	if err := decoder.Decode(&result); err != nil {
		panic(err)
	}

	// convert struct to json and dump the output
	mapJson, _ := json.Marshal(result)
    log.Println(string(mapJson))
}
