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
var doneCh = make(chan struct{}) //struct with no fiels
//reequires 0 memory allocation
//to jest signal only channel

func main() {
	go logger()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{}
}

func logger() {
	for {
		select { //blokuje do czasu aż któryś case otrzyma wiadomość
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n",
				entry.time.Format("2006-01-02T15:04:05"),
				entry.severity,
				entry.message)
		case <-doneCh:
			break
		} //nie może być defaulta, jakby był to select nie będzie blokujący
	}
}
