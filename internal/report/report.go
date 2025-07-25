package report

import (
	"fmt"

	"github.com/mobenaus/fc-pos-go-desafio-stress-test/internal/stress"
)

func DisplayReport(str *stress.StressTestResults) {
	fmt.Printf("Stress Test executado em %v\n", str.TotalTime)
	fmt.Printf("Total requests........: %v\n", str.TotalRequests)
	fmt.Printf("Total Success requests: %v\n", str.Total200Requests)
	str.TotalStatusMap.Range(func(code, total any) bool {
		if totalInt64, ok := total.(*int64); ok {
			fmt.Printf("Total %v requests....: %v\n", code, *totalInt64)
		}
		return true
	})
	fmt.Printf("Requests por segundo %f\n", float64(str.TotalRequests)/str.TotalTime.Seconds())

}
