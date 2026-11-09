package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	log := NewLog()

	if len(os.Args) < 2 {
		fmt.Println("Expected portnumber as argument")
		os.Exit(1)
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("invalid portnumber: %v\n", os.Args[1])
		os.Exit(1)
	}

	running_port := SysInfo{
		Msg:       fmt.Sprintf("Running server on: %d", port),
		Timestamp: time.Now(),
		Level:     "INFO",
	}
	log.Event = append(log.Event, running_port)

	running_port.PrintInfo()

}

// Start the HTTP server
func startServer(args []string, port int) {

	// Create a router to delegate requests to the server
	r := router.NewRouter(issueService)

	fs.Parse(args)

	// set up the HTTP server
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080", // Add configuration if want to change ports etc.
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
	log.Fatal(server.ListenAndServe())
}

type Log struct {
	Event []SysInfo
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
	fmt.Printf("[%s] %v - %s", s.Level, s.Msg, s.Timestamp)
}
