package simulation

import (
	"fmt"
	"time"

	"github.com/rafaquelhodev/hammer/job"
)

// Struct that defines the orchestrator
type Orchestrator struct {
	// Interval between execution calls
	Interval time.Duration
	// Total duration the orchestrator must run
	TotalDuration time.Duration
}

// Starts the orchestrator
func (o *Orchestrator) Start(job job.Job) {
	ticker := time.NewTicker(o.Interval)
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

	time.Sleep(o.TotalDuration)
	ticker.Stop()

	done <- true
	fmt.Println("Done!!")
}
