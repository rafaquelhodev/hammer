package tester_test

import (
	"testing"
	"time"

	"github.com/rafaquelhodev/hammer/config"
	"github.com/rafaquelhodev/hammer/job"
	"github.com/rafaquelhodev/hammer/tester"
)

func TestRunTester(t *testing.T) {
	start := time.Now()

	config := config.Config{DomainUrl: "test.com", TimeSettings: config.TimeSettings{Frequency: time.Millisecond * 100, TotalDuration: time.Second * 1}}

	job := job.HttpJob{DomainUrl: config.DomainUrl}

	tester.Start(job, config)

	if elapsed := time.Since(start); elapsed.Seconds() < 1.0 {
		t.Fatalf("The elapsed time was less than 1 second")
	}

}
