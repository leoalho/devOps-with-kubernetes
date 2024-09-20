package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	appUUID := uuid.New()

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			timestamp := time.Now().Format(time.RFC3339)

			fmt.Printf("%s %s\n", timestamp, appUUID)
		}
	}
}
