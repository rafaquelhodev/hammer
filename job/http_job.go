package job

import (
	"fmt"
	"time"
)

type HttpJob struct {
	DomainUrl string
}

func (job HttpJob) Execute() Response {
	fmt.Println("Executing job")
	return Response{Status: "200", RunDuration: time.Millisecond * 200, Success: true}
}
