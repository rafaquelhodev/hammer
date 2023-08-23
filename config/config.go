package config

import "time"

type Config struct {
	DomainUrl    string
	TimeSettings TimeSettings
}

// Struct for the time settings when running hammer
type TimeSettings struct {
	// Frequency between execution calls
	Frequency time.Duration
	// Total duration hammer must run
	TotalDuration time.Duration
}
