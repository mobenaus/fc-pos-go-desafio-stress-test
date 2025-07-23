package report

import (
	"fmt"

	"github.com/mobenaus/fc-pos-go-desafio-stress-test/internal/stress"
)

func DisplayReport(str stress.StressTestResults) {
	fmt.Printf("Stress Test executado em %v\n", str.TotalTime)

}
