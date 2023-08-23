package main

import (
	"time"

	"github.com/rafaquelhodev/hammer/config"
	"github.com/rafaquelhodev/hammer/job"
	"github.com/rafaquelhodev/hammer/tester"
)

func main() {
	config := config.Config{DomainUrl: "test.com", TimeSettings: config.TimeSettings{Frequency: time.Millisecond * 300, TotalDuration: time.Second * 2}}

	job := job.HttpJob{DomainUrl: config.DomainUrl}

	tester.Start(job, config)
}
