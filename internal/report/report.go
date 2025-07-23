package report

import (
	"fmt"

	"github.com/mobenaus/fc-pos-go-desafio-stress-test/internal/stress"
)

func DisplayReport(str stress.StressTestResults) {
	fmt.Printf("Stress Test executado em %v\n", str.TotalTime)
	fmt.Printf("Total requests........: %v\n", str.TotalRequests)
	fmt.Printf("Total Success requests: %v\n", str.Total200Requests)
	for code, total := range str.TotalStatusMap {
		fmt.Printf("Total %v requests....: %v\n", code, total)
	}
}
