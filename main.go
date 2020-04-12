package main

import (
	"fmt"
	"time"
)

func work() {
	c := 0.0
	for i := 1; i < 100000; i++ {
		a := 1.1
		b := 2.2
		c += a / b
	}
	//time.Sleep(time.Second)
}

func main() {

	type snapshot struct {
		time            time.Time
		qpc             time.Duration
		unbiased        time.Duration
		unbiasedPrecise time.Duration
	}

	ss := make([]snapshot, 10)
	for i := range ss {
		work()
		ss[i].time = time.Now()
		ss[i].qpc = QPC()
		ss[i].unbiased = UnbiasedInterruptTime()
		ss[i].unbiasedPrecise = UnbiasedInterruptPreciseTime()
	}

	for i := range ss {
		if i == 0 {
			continue
		}
		fmt.Printf("time=%v", ss[i].time.Sub(ss[i-1].time))
		fmt.Printf("\tqpc=%v", ss[i].qpc-ss[i-1].qpc)
		fmt.Printf("\tunbiased=%v", ss[i].unbiased-ss[i-1].unbiased)
		fmt.Printf("\tunbiasedPrecise=%v", ss[i].unbiasedPrecise-ss[i-1].unbiasedPrecise)
		fmt.Printf("\n")
	}
}
