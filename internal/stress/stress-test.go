package stress

import (
	"context"
	"time"
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
	start := time.Now()

	time.Sleep(10 * time.Second)

	finish := time.Now()
	return StressTestResults{
		TotalTime: finish.Sub(start),
	}, nil
}
