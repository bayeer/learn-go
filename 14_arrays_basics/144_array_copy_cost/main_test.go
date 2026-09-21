package main

import (
	"strconv"
	"testing"
)

func BenchmarkProcessByVal(b *testing.B) {
	large := LargeArr{}
	for b.Loop() {
		processByVal(large)
	}
}

func BenchmarkProcessByPtr(b *testing.B) {
	large := LargeArr{}
	for b.Loop() {
		processByPtr(&large)
	}
}

func BenchmarkSubBenchmarks(b *testing.B) {
	inputs := [...]int{200, 2000, 20000, 200000}
	large := LargeArr{}
	for _, size := range inputs {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			for b.Loop() {
				processByVal(large)
			}
		})
	}
	for _, size := range inputs {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			for b.Loop() {
				processByPtr(&large)
			}
		})
	}
}
