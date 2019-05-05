package main

import (
	"testing"
	"time"
)

var (
	t time.Time
	d time.Duration
)

func BenchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t = time.Now()
	}
}

func BenchmarkQPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d = QPC()
	}
}

func BenchmarkUnbiasedInterruptTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d = UnbiasedInterruptTime()
	}
}

func BenchmarkUnbiasedInterruptPreciseTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d = UnbiasedInterruptPreciseTime()
	}
}
