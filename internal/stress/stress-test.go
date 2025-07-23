package stress

import (
	"context"
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
	TotalRequests    int
	Total200Requests int
	TotalStatusMap   map[int]int
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

func (st *StressTest) Execute(ctx context.Context) (StressTestResults, error) {

	result := StressTestResults{
		TotalStatusMap: make(map[int]int),
	}
	request, error := request.NewStressRequest(st.url)
	if error != nil {
		return StressTestResults{}, error
	}

	start := time.Now()

	result.TotalRequests++

	status, error := request.Execute()
	if error != nil {
		return StressTestResults{}, error
	}

	if status == 200 {
		result.Total200Requests++
	}
	result.TotalStatusMap[status]++

	finish := time.Now()

	result.TotalTime = finish.Sub(start)

	return result, nil
}
