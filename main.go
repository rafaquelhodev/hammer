package main

import (
	"time"

	"github.com/rafaquelhodev/hammer/job"
	"github.com/rafaquelhodev/hammer/simulation"
)

func main() {
	job := job.HttpJob{DomainUrl: "test.com"}

	orchestrator := simulation.Orchestrator{Interval: time.Millisecond * 300, TotalDuration: time.Second * 2}
	orchestrator.Start(job)
}
