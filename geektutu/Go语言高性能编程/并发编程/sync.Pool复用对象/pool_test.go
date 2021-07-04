/*
sync.Pool 的使用场景
	保存和复用临时对象，减少内存分配，降低 GC 压力
 */

package main

import (
	"encoding/json"
	"sync"
	"testing"
)

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

/*
go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: zwd.com/geektutu/Go语言高性能编程/并发编程/sync.Pool复用对象
BenchmarkUnmarshal-4                8576            147077 ns/op            1384 B/op          7 allocs/op
BenchmarkUnmarshalWithPool-4        9223            142606 ns/op             232 B/op          6 allocs/op
PASS
ok      zwd.com/geektutu/Go语言高性能编程/并发编程/sync.Pool复用对象    3.557s
 */