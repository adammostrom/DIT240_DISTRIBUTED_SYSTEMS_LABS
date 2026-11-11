package main

import (
	"fmt"
	"log"
	"net" // https://pkg.go.dev/net
	"os"
	"strconv"
)

const LOCALHOST = "https://localhost/"
const LOCALHOST_127 = "127.0.0.1"

func main() {

	log := NewLog()

	if len(os.Args) < 2 {
		fmt.Println("Expected portnumber as argument")
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("invalid portnumber: %v. Error: %v\n", os.Args[1], err)
		os.Exit(1)
	}

	if !validatePort(port) {
		fmt.Printf("Port number out of bounds. Provided: %d, range: %d - %d\n", port, ALLOWED_PORT_LOWER, ALLOWED_PORT_UPPER)
		os.Exit(1)
	}

	//connectToServer()
	createServer(port, LOCALHOST_127, log)

}

// Connect to a server:

func connectToServer(port int, host string) {
	addr := net.JoinHostPort(host, strconv.Itoa(port))

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Write a message to the server. "_" = bytecount
	_, err = conn.Write([]byte("Hello from client"))
	if err != nil {
		log.Fatal()
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
}

// Create a server
func createServer(port int, host string, log *Log) error {

	addr := net.JoinHostPort(host, strconv.Itoa(port))

	// Assume tcp
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.LogEvent(fmt.Sprintf("listen on %s: %v", addr, err), LevelError)
		return fmt.Errorf("listen on %s: %w", addr, err)
	}
	log.LogEvent(fmt.Sprintf("Running server on port: %d", port), LevelInfo)
	defer ln.Close()
	for {
		// conn = a client
		conn, err := ln.Accept()
		if err != nil {
			log.LogEvent(fmt.Sprintf("error: %v", err), LevelError)
			continue
		}

		go handleConnection(conn, addr, log)
	}
}

// Active connection
func handleConnection(conn net.Conn, addr string, log *Log) {

	defer conn.Close()
	fmt.Println("Accepted connection")
	buf := make([]byte, 1024)

	// Expect client to write something
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.LogEvent(fmt.Sprintf("read error: %v\n", err), LevelError)
			fmt.Printf("read error: %v\n", err)
			return
		}
		request := string(buf[:n])
		log.LogEvent(fmt.Sprintf("received: %s\n", request), LevelInfo)
		fmt.Printf("received: %s\n", request)
		_, err = conn.Write([]byte("Hello from server!\n"))
		if err != nil {
			fmt.Printf("write error: %v\n", err)
			return
		}
	}

}

// Start the HTTP server
func startServer(args []string, port int) {

	// set up the HTTP server
	/*
		port_str := strconv.Itoa(port)

		server := &http.Server{
			Addr:    fmt.Sprintf(":%s", port_str),
			Handler: mux,
		}

		// Delegates all HTTP requests to /issues* to the IssuesHandler
		mux.HandleFunc("/issues", r.AllRouting)  // for /issues exact (list, create)
		mux.HandleFunc("/issues/", r.AllRouting) // for /issues/{id} (single issue)

		log.Println("Connected successfully — server starting...")

		// Serve the HTML frontend
		mux.Handle("/", http.FileServer(http.Dir("./static")))

		// Keep server running on port 8080
		log.Println("Server running at http://localhost:8080")
		log.Fatal(server.ListenAndServe()) */
}
