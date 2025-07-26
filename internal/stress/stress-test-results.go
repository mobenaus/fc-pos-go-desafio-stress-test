package stress

import (
	"time"
)

type StressTestResults struct {
	TotalTime        time.Duration
	TotalRequests    int64
	Total200Requests int64
	TotalStatusMap   map[int]int64
	start            time.Time
}

func NewStressTestResults() *StressTestResults {
	return &StressTestResults{
		TotalStatusMap: make(map[int]int64),
	}
}

func (str *StressTestResults) Start() {
	str.start = time.Now()
}

func (str *StressTestResults) Finish() {
	str.TotalTime = time.Since(str.start)
}

func (str *StressTestResults) CountResult(status int) {
	str.TotalRequests++

	if status == 200 {
		str.Total200Requests++
	} else {
		str.TotalStatusMap[status]++
	}

}
