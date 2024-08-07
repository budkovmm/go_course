package main

import (
	"testing"
)

func BenchmarkPushFront(b *testing.B) {
	d := NewDeque()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.PushFront(i)
	}
}
