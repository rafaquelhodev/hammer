package tester

import (
	"fmt"
	"time"

	"github.com/rafaquelhodev/hammer/config"
	"github.com/rafaquelhodev/hammer/job"
)

// Starts the hammer test
func Start(hammer job.Job, config config.Config) {
	ticker := time.NewTicker(config.TimeSettings.Frequency)
	done := make(chan bool)
	hammerResultsChannel := make(chan [][]job.Response)
	defer close(hammerResultsChannel)

	go func() {
		results := [][]job.Response{}
		for {
			select {
			case <-done:
				hammerResultsChannel <- results
				fmt.Print("Closing Jobchannel \n")
				return
			case t := <-ticker.C:
				response := hammer.Execute()
				results = append(results, response)
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(config.TimeSettings.TotalDuration)
	ticker.Stop()

	done <- true

	hammerResults := <-hammerResultsChannel
	fmt.Printf("Hammer results %v \n", hammerResults)

	fmt.Println("Done!!")
}
