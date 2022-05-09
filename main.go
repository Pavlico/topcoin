package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func getSIGRTchannel() chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	sigArr := make([]os.Signal, 31)
	for i := range sigArr {
		sigArr[i] = syscall.Signal(i + 0x22)
	}
	log.Println(sigArr)

	signal.Notify(sigChan, sigArr...)
	return sigChan
}

func main() {
	c := getSIGRTchannel()
	// Block until a signal is received.
	for {
		s := <-c
		fmt.Println("Got signal:", s)
	}
}
