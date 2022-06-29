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

//kiedy tworzymy gorutyny i nie ma oczywistej informacji
//kiedy się zakończą - o tym jest ten topic
var logCh = make(chan logEntry, 50)

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Millisecond)
	//i tutaj bo main się kończy channel i gorutyna są ubijane
	//czasami to wystarczy ale ogólnie słabo
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v]%v\n",
			entry.time.Format("2006-01-02T15:04:05"),
			entry.severity,
			entry.message)
	}
}
