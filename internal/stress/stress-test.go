package stress

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mobenaus/fc-pos-go-desafio-stress-test/internal/request"
)

type StressTest struct {
	url         string
	requests    int
	concurrency int
}

func NewStressTest(
	url string,
	requests int,
	concurrency int) *StressTest {
	return &StressTest{
		url:         url,
		requests:    requests,
		concurrency: concurrency,
	}
}

func (st *StressTest) Execute(ctx context.Context) (*StressTestResults, error) {

	result := NewStressTestResults()

	request, error := request.NewStressRequest(st.url)
	if error != nil {
		return result, error
	}

	result.Start()
	defer result.Finish()

	requestchannel := make(chan int)
	resultchannel := make(chan int, st.requests)
	var requestwg sync.WaitGroup
	var resultwg sync.WaitGroup

	go func() {
		for {
			status := <-resultchannel
			result.CountResult(status)
			resultwg.Done()
		}
	}()

	for range st.concurrency {
		go func() {
			for {
				<-requestchannel
				status, err := request.Execute()
				if err != nil {
					fmt.Println(err.Error())
				}
				resultwg.Add(1)
				resultchannel <- status
				requestwg.Done()
			}
		}()
	}

	var i int

	go reportStatusRoutine(&i)

	for i = 0; i < st.requests; i++ {
		requestwg.Add(1)
		requestchannel <- i
	}

	requestwg.Wait()
	resultwg.Wait()

	return result, nil
}

func reportStatusRoutine(counter *int) {
	for {
		time.Sleep(time.Second)
		fmt.Printf("Called %d requests...\r", *counter)
	}
}
