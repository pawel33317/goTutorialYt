package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

//tu jest poprawka, żeby zamknąć channel i gorutyne
//i mamy gracefull shutdown
var logCh = make(chan logEntry, 50)

func main() {
	go logger()
	defer func() {
		close(logCh)
	}()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n",
			entry.time.Format("2006-01-02T15:04:05"),
			entry.severity,
			entry.message)
	}
}
