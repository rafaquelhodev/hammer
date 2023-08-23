package job

import (
	"fmt"
	"time"
)

type HttpJob struct {
	DomainUrl    string
	HammerNumber int32
}

func (job HttpJob) Execute() []Response {
	hammerChannel := make(chan Response, job.HammerNumber)
	defer close(hammerChannel)

	for hammer := 1; hammer <= int(job.HammerNumber); hammer++ {
		go func(hammer int) {
			fmt.Printf("Executing job %v \n", hammer)
			hammerChannel <- Response{Status: "200", RunDuration: time.Millisecond * 200, Success: true}
		}(hammer)
	}

	var responses []Response
	for hammer := 1; hammer <= int(job.HammerNumber); hammer++ {
		response := <-hammerChannel
		responses = append(responses, response)
	}

	return responses
}
