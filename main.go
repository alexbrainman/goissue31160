package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	startQPC := QPC()
	startUnbiased := UnbiasedInterruptTime()
	startUnbiasedPrecise := UnbiasedInterruptPreciseTime()
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second)

		nowTime := time.Now()
		elapsedTime := nowTime.Sub(startTime)

		fmt.Printf("time=%v ", elapsedTime)

		nowQPC := QPC()
		elapsedQPC := nowQPC - startQPC

		fmt.Printf("qpc=( %v / %v ) ", elapsedQPC, elapsedTime-elapsedQPC)

		nowUnbiased := UnbiasedInterruptTime()
		elapsedUnbiased := nowUnbiased - startUnbiased

		fmt.Printf("unbiased=( %v / %v ) ", elapsedUnbiased, elapsedTime-elapsedUnbiased)

		nowUnbiasedPrecise := UnbiasedInterruptPreciseTime()
		elapsedUnbiasedPrecise := nowUnbiasedPrecise - startUnbiasedPrecise

		fmt.Printf("unbiased_precise=( %v / %v ) ", elapsedUnbiasedPrecise, elapsedTime-elapsedUnbiasedPrecise)

		fmt.Println()
	}

}
