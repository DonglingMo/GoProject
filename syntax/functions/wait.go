package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not ready, retrying... (%d/%d)\n", tries, int(timeout.Seconds()))
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond within %s", url, timeout)
}
