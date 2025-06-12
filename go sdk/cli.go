package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Go SDK CLI Tool - Error Listener (CLI Demo)")

	errorChan := make(chan error)
	ListenForErrors(errorChan, SendEmailNotification)

	// currently we simulate errors but we here we wanna replace it with a real error
	// like the one that actively watches teminal logs
	go func() {
		errorChan <- fmt.Errorf("simulated error: something went wrong")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Shutting down error listener.")
} 