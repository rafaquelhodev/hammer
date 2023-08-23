package main

import (
	"time"

	"github.com/rafaquelhodev/hammer/config"
	"github.com/rafaquelhodev/hammer/job"
	"github.com/rafaquelhodev/hammer/tester"
)

func main() {
	config := config.Config{DomainUrl: "test.com", TimeSettings: config.TimeSettings{Frequency: time.Millisecond * 3000, TotalDuration: time.Second * 6}}

	job := job.HttpJob{DomainUrl: config.DomainUrl, HammerNumber: 2}

	tester.Start(job, config)
}
