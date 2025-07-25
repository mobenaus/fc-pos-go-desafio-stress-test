package stress

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mobenaus/fc-pos-go-desafio-stress-test/internal/request"
)

type StressTest struct {
	url         string
	requests    int
	concurrency int
}

type StressTestResults struct {
	TotalTime        time.Duration
	TotalRequests    int64
	Total200Requests int64
	TotalStatusMap   sync.Map
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

	result := &StressTestResults{
		TotalStatusMap: sync.Map{},
	}
	request, error := request.NewStressRequest(st.url)
	if error != nil {
		return result, error
	}

	start := time.Now()

	defer func() {
		result.TotalTime = time.Since(start)
	}()

	rc := make(chan int)
	var wg sync.WaitGroup

	for range st.concurrency {
		go func() {
			for {
				<-rc
				err := st.callRequest(result, request)
				if err != nil {
					fmt.Println(err.Error())
				}
				wg.Done()
			}
		}()
	}

	var i int

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("Called %d requests...\r", i)
		}
	}()

	for i = 0; i < st.requests; i++ {
		wg.Add(1)
		rc <- i
	}
	wg.Wait()
	return result, nil
}

func (st *StressTest) callRequest(result *StressTestResults, request *request.StressRequest) error {
	atomic.AddInt64(&result.TotalRequests, 1)

	status, error := request.Execute()
	if error != nil {
		return error
	}

	if status == 200 {
		atomic.AddInt64(&result.Total200Requests, 1)
	} else {
		val, _ := result.TotalStatusMap.LoadOrStore(status, new(int64))
		ptr := val.(*int64)
		atomic.AddInt64(ptr, 1)
	}
	return nil
}
