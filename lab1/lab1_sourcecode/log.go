package main

import (
	"fmt"
	"sync"
	"time"
)

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
