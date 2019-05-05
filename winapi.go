package main

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	modkernel32                       = syscall.NewLazyDLL("kernel32.dll")
	queryPerformanceFrequency         = modkernel32.NewProc("QueryPerformanceFrequency")
	queryPerformanceCounter           = modkernel32.NewProc("QueryPerformanceCounter")
	queryUnbiasedInterruptTime        = modkernel32.NewProc("QueryUnbiasedInterruptTime")
	modkernelbase                     = syscall.NewLazyDLL("kernelbase.dll")
	queryUnbiasedInterruptTimePrecise = modkernelbase.NewProc("QueryUnbiasedInterruptTimePrecise")
)

var (
	qpcStartCounter int64
	qpcFrequency    int64
)

func init() {
	syscall.Syscall(queryPerformanceCounter.Addr(), 1, uintptr(unsafe.Pointer(&qpcStartCounter)), 0, 0)
	r1, _, _ := syscall.Syscall(queryPerformanceFrequency.Addr(), 1, uintptr(unsafe.Pointer(&qpcFrequency)), 0, 0)
	if r1 == 0 {
		panic("QueryPerformanceFrequency failed")
	}
}

func QPC() time.Duration {
	var c int64 = 0
	syscall.Syscall(queryPerformanceCounter.Addr(), 1, uintptr(unsafe.Pointer(&c)), 0, 0)
	// TODO: rewrite this expression, because it is overflow on my PC after about 15 minutes
	return time.Duration((c - qpcStartCounter) * 1000000000 / qpcFrequency)
}

func UnbiasedInterruptTime() time.Duration {
	var c uint64 = 0
	syscall.Syscall(queryUnbiasedInterruptTime.Addr(), 1, uintptr(unsafe.Pointer(&c)), 0, 0)
	return time.Duration(c) * 100
}

func UnbiasedInterruptPreciseTime() time.Duration {
	var c uint64 = 0
	syscall.Syscall(queryUnbiasedInterruptTimePrecise.Addr(), 1, uintptr(unsafe.Pointer(&c)), 0, 0)
	return time.Duration(c) * 100
}
