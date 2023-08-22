package config

import "time"

type Config struct {
	DomainUrl    string
	TimeSettings TimeSettings
}

// Struct that defines the orchestrator
type TimeSettings struct {
	// Interval between execution calls
	Interval time.Duration
	// Total duration the orchestrator must run
	TotalDuration time.Duration
}
