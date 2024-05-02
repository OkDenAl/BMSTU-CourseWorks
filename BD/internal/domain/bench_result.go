package domain

import (
	"fmt"
	"time"
)

type BenchResult struct {
	Name            string
	NumOfOperations int
	UsedTime        time.Duration
}

func NewBenchResult(name string, numOfOperations int) *BenchResult {
	return &BenchResult{Name: name, NumOfOperations: numOfOperations}
}

func (br BenchResult) String() string {
	return fmt.Sprintf("Benchmark name: %s\nNumber of operations: %d\nUsed time: %s\n",
		br.Name, br.NumOfOperations, br.UsedTime.String())
}
