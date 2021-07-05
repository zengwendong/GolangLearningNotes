package main

import (
	"sync"
	"testing"
	"time"
)

type RW interface {
	Write()
	Read()
}

const cost = time.Microsecond

// 采用互斥锁
type Lock struct {
	count int
	mu    sync.Mutex
}

func (l *Lock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *Lock) Read() {
	l.mu.Lock()
	time.Sleep(cost)
	_ = l.count
	l.mu.Unlock()
}

// 采用读写锁
type RWLock struct {
	count int
	mu    sync.RWMutex
}

func (l *RWLock) Write() {
	l.mu.Lock()
	l.count++
	time.Sleep(cost)
	l.mu.Unlock()
}

func (l *RWLock) Read() {
	l.mu.RLock()
	_ = l.count
	time.Sleep(cost)
	l.mu.RUnlock()
}

func benchmark(b *testing.B, rw RW, read, write int) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < read*100; k++ {
			wg.Add(1)
			go func() {
				rw.Read()
				wg.Done()
			}()
		}
		for k := 0; k < write*100; k++ {
			wg.Add(1)
			go func() {
				rw.Write()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}


func BenchmarkReadMore(b *testing.B)    { benchmark(b, &Lock{}, 9, 1) }
func BenchmarkReadMoreRW(b *testing.B)  { benchmark(b, &RWLock{}, 9, 1) }
func BenchmarkWriteMore(b *testing.B)   { benchmark(b, &Lock{}, 1, 9) }
func BenchmarkWriteMoreRW(b *testing.B) { benchmark(b, &RWLock{}, 1, 9) }
func BenchmarkEqual(b *testing.B)       { benchmark(b, &Lock{}, 5, 5) }
func BenchmarkEqualRW(b *testing.B)     { benchmark(b, &RWLock{}, 5, 5) }

/*
go test -bench .
goos: windows
goarch: amd64
pkg: zwd.com/geektutu/Go语言高性能编程/并发编程/读写锁和互斥锁的性能比较
BenchmarkReadMore-4                    1        1006453200 ns/op
BenchmarkReadMoreRW-4                 10         102968750 ns/op
BenchmarkWriteMore-4                   1        1002429200 ns/op
BenchmarkWriteMoreRW-4                 2         905534300 ns/op
BenchmarkEqual-4                       2        1004061950 ns/op
BenchmarkEqualRW-4                     2         502739450 ns/op
 */

/*
读写锁的存在是为了解决读多写少时的性能问题，读场景较多时，读写锁可有效地减少锁阻塞的时间

读锁之间不互斥，没有写锁的情况下，读锁是无阻塞的，多个协程可以同时获得读锁。
写锁之间是互斥的，存在写锁，其他写锁阻塞。
写锁与读锁是互斥的，如果存在读锁，写锁阻塞，如果存在写锁，读锁阻塞。

Go 标准库中提供了 sync.RWMutex 互斥锁类型及其四个方法：
Lock 加写锁
Unlock 释放写锁
RLock 加读锁
RUnlock 释放读锁
 */