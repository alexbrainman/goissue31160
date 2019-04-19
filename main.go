package main

import (
	"fmt"
	"sort"
	"syscall"
	"time"
	"unsafe"
	_ "unsafe"
)

// precision timing
var (
	modkernel32                = syscall.NewLazyDLL("kernel32.dll")
	_QueryPerformanceFrequency = modkernel32.NewProc("QueryPerformanceFrequency")
	_QueryPerformanceCounter   = modkernel32.NewProc("QueryPerformanceCounter")
)

// now returns time.Duration using queryPerformanceCounter
func QPC() int64 {
	var now int64
	syscall.Syscall(_QueryPerformanceCounter.Addr(), 1, uintptr(unsafe.Pointer(&now)), 0, 0)
	return now
}

// QPCFrequency returns frequency in ticks per second
func QPCFrequency() int64 {
	var freq int64
	r1, _, _ := syscall.Syscall(_QueryPerformanceFrequency.Addr(), 1, uintptr(unsafe.Pointer(&freq)), 0, 0)
	if r1 == 0 {
		panic("call failed")
	}
	return freq
}

const N = 1e7

func main() {
	qpcFrequency := QPCFrequency()
	timeSamples := make([]time.Time, N)
	qpcSamples := make([]int64, N)

	for i := range qpcSamples {
		qpcSamples[i] = QPC()
	}

	for i := range timeSamples {
		timeSamples[i] = time.Now()
	}

	timeDeltas := make([]int64, N-1)
	qpcDeltas := make([]int64, N-1)

	for i, next := range timeSamples[1:] {
		timeDeltas[i] = next.Sub(timeSamples[i]).Nanoseconds()
	}

	for i, next := range qpcSamples[1:] {
		qpcDeltas[i] = (next - qpcSamples[i]) * 1e9 / qpcFrequency
	}

	sort.Slice(timeDeltas, func(i, k int) bool { return timeDeltas[i] < timeDeltas[k] })
	sort.Slice(qpcDeltas, func(i, k int) bool { return qpcDeltas[i] < qpcDeltas[k] })

	timeZeros := 0
	for i, v := range timeDeltas {
		if v == 0 {
			timeZeros++
		} else {
			timeDeltas = timeDeltas[i:]
			break
		}
	}

	qpcZeros := 0
	for i, v := range qpcDeltas {
		if v == 0 {
			qpcZeros++
		} else {
			qpcDeltas = qpcDeltas[i:]
			break
		}
	}

	time50 := timeDeltas[len(timeDeltas)*50/100]
	time90 := timeDeltas[len(timeDeltas)*90/100]
	time99 := timeDeltas[len(timeDeltas)*99/100]

	qpc50 := qpcDeltas[len(qpcDeltas)*50/100]
	qpc90 := qpcDeltas[len(qpcDeltas)*90/100]
	qpc99 := qpcDeltas[len(qpcDeltas)*99/100]

	fmt.Printf("              %-8v  %8v  %8v  %8v\n", "", "50%", "90%", "99%")
	fmt.Printf("time.Now    Z=%-8v  %8v  %8v  %8v\n", timeZeros, time50, time90, time99)
	fmt.Printf("QPC         Z=%-8v  %8v  %8v  %8v\n", qpcZeros, qpc50, qpc90, qpc99)
}
