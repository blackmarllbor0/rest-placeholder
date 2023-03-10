package signal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ListenSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sigVal := <-sigChan
	signal.Stop(sigChan)
	log.Println("stop signal", sigVal)
}
