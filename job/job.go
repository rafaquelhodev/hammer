package job

import "time"

type Job interface {
	Execute() []Response
}

type Response struct {
	Status      string
	Success     bool
	RunDuration time.Duration
}
