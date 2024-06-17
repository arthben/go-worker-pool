package main

import (
	"time"
)

type EmailTask struct {
	EmailTo     string
	Subject     string
	MessageBody string
}

// Proccess Email Task
func (t *EmailTask) Proccess() string {
	// this is the function that implement Process from the interface Task
	time.Sleep(2 * time.Second)
	return "Send Email To " + t.EmailTo
}
