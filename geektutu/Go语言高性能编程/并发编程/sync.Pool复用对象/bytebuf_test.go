package main

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		// A Buffer is a variable-sized buffer of bytes with Read and Write methods
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		// Write appends the contents of p to the buffer, growing the buffer as needed
		buf.Write(data)
		// Reset resets the buffer to be empty
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}

/*
BenchmarkBufferWithPool-4        6352879               176 ns/op               0 B/op          0 allocs/op
BenchmarkBuffer-4                 449824              2355 ns/op           10240 B/op          1 allocs/op
 */