package main

import (
	"bufio"
	"fmt"
	"net" // https://pkg.go.dev/net
	"os"
	"strconv"
	"sync"
	"time"
)

const LOCALHOST = "http://localhost/"

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
	createServer(port, LOCALHOST, log)
	running_port := SysInfo{
		Msg:       fmt.Sprintf("Running server on port: %d", port),
		Timestamp: time.Now(),
		Level:     "INFO",
	}
	log.Event = append(log.Event, running_port)

	running_port.PrintInfo()

}

// Connect to a server:

func connectToServer() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Printf("status: %v\n", status)
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
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.LogEvent(fmt.Sprintf("error: %v", err), LevelError)
			return fmt.Errorf("error: %w", err)
		}
		go handleConnection(conn, addr, log)
	}
}

// Active connection
func handleConnection(conn net.Conn, addr string, log *Log) {

	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		log.LogEvent(fmt.Sprintf("read error: %v\n", err), LevelError)
		fmt.Printf("read error: %v\n", err)
		return
	}
	request := string(buf[:n])
	log.LogEvent(fmt.Sprintf("received: %s\n", request), LevelInfo)
	fmt.Printf("received: %s\n", request)

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

type Log struct {
	mu    sync.Mutex
	Event []SysInfo
}

const (
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
)

// Appends a SysInfo to the Log and prints the SysInfo using fmt.Printf
func (l *Log) LogEvent(msg string, level string) {

	sysInfo := SysInfo{
		Msg:       msg,
		Timestamp: time.Now(),
		Level:     level,
	}
	l.mu.Lock()
	l.Event = append(l.Event, sysInfo)
	l.mu.Unlock()

	// Read only, no need for sync
	sysInfo.PrintInfo()

}
func NewLog() *Log {
	return &Log{}
}

// Expand with more parameters as it grows
type SysInfo struct {
	Msg       string
	Timestamp time.Time
	Level     string
}

func (s *SysInfo) PrintInfo() {
	fmt.Printf("[%s] %s  %v \n", s.Level, s.Timestamp.Format("15:04:05"), s.Msg)
}

// Allowed ports
const ALLOWED_PORT_UPPER = 65355
const ALLOWED_PORT_LOWER = 1024

func validatePort(port int) bool {
	if port > ALLOWED_PORT_LOWER && port < ALLOWED_PORT_UPPER {
		return true
	}
	return false
}
