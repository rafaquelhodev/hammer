package tester

import (
	"fmt"
	"time"

	"github.com/rafaquelhodev/hammer/config"
	"github.com/rafaquelhodev/hammer/job"
)

// Starts the hammer test
func Start(job job.Job, config config.Config) {
	ticker := time.NewTicker(config.TimeSettings.Frequency)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				job.Execute()
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(config.TimeSettings.TotalDuration)
	ticker.Stop()

	done <- true
	fmt.Println("Done!!")
}
