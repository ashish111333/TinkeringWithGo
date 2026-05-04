package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/ashish111333/twgo/exercises"
)

var s []int64 = RandIntSlice(8000000, 100, 8000000, false)

func TestAddSliceItems(t *testing.T) {
	if AddSLiceItems(s) != (AddSliceItemsC(12, s)) {
		t.Fatalf("outputs don't match")
	}
}
func TestShardedCounter(t *testing.T) {

	sc := NewShardedCounter(6)

	var countUpTo int64 = 100
	for i := range countUpTo {
		sc.Inc(i)
	}

	if sc.Total() != countUpTo {
		fmt.Println(sc.Total())
		t.FailNow()
	}

}
func TestSimpleCounter(t *testing.T) {
	var countUpto int64 = 100

	nc := NewCounter(countUpto)
	for range countUpto {
		nc.Inc()
	}
	if nc.Count != 100 {
		t.FailNow()
	}
}

func BenchmarkShardedCounterVsConcurrentcounter(b *testing.B) {
	const (
		shards           int64 = 64
		incrementsPerGor int64 = 5000
	)
	workers := int64(runtime.GOMAXPROCS(0) * 4)

	b.Run("sharded_counter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sc := NewShardedCounter(shards)
			var wg sync.WaitGroup
			wg.Add(int(workers))
			for workerID := range workers {
				go func(workerID int64) {
					defer wg.Done()
					for range incrementsPerGor {
						sc.Inc(workerID)
					}
				}(workerID)
			}
			wg.Wait()

			want := workers * incrementsPerGor
			if got := sc.Total(); got != want {
				b.Fatalf("sharded total = %d, want %d", got, want)
			}
		}
	})

	b.Run("normal_counter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nc := NewCounter(workers * incrementsPerGor)
			var wg sync.WaitGroup
			wg.Add(int(workers))
			for range workers {
				go func() {
					defer wg.Done()
					for range incrementsPerGor {
						nc.Inc()
					}
				}()
			}
			wg.Wait()

			want := workers * incrementsPerGor
			if got := nc.Count; got != want {
				b.Fatalf("normal total = %d, want %d", got, want)
			}
		}
	})
}

func BenchmarkAddSliceItems(b *testing.B) {
	var sC int64
	b.Run("AddSLiceItemsC", func(b *testing.B) {
		for b.Loop() {
			sC = AddSliceItemsC(8000, s)
		}
		b.Logf("sum --->%d", sC)
	})
	var sCC int64
	b.Run("AddSliceItemsCButChannels", func(b *testing.B) {
		for b.Loop() {
			sCC = addSliceItemsCChannels(8000, s)
		}
		b.Logf("sum given by channels one-->%d", sCC)
	})
	var sCMx int64
	b.Run("AddSliceItemsCMx", func(b *testing.B) {

		for b.Loop() {
			sCMx = AddSliceItemsCMx(8000, s)
		}
		b.Logf("sum given by mutex approach--->%d", sCMx)
	})

	b.Run("AddSliceItems", func(b *testing.B) {
		for b.Loop() {
			AddSLiceItems(s)
		}
	})

}

func BenchmarkUpdateVar(b *testing.B) {
	a := 0
	times := 500
	b.Run("updateVar", func(b *testing.B) {
		UpdateVar(&a, times)
	})
	e := 0
	b.Run("UpdateVarCh", func(b *testing.B) {
		UpdateVarCh(&e, times)
	})
	c := 0
	b.Run("UpdateVarMx", func(b *testing.B) {
		UpdateVarMx(&c, times)
	})
}

func BenchmarkCounterAtomicsVsMx(b *testing.B) {
	var num1 int64 = 0
	var num2 int64 = 0

	incBy := 100000000
	numGoRoutines := incBy
	b.Run("counterUsingAtomics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			wg.Add(numGoRoutines)
			for range numGoRoutines {
				go func() {
					defer wg.Done()
					atomic.AddInt64(&num1, 1)
				}()

			}
			wg.Wait()

		}

	})
	b.Run("counterWithMx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var mx sync.Mutex
			var wg sync.WaitGroup
			wg.Add(numGoRoutines)
			for range numGoRoutines {
				go func() {
					mx.Lock()
					defer mx.Unlock()
					defer wg.Done()
					num2 += 1
				}()
			}
			wg.Wait()

		}
	})

}

func BenchmarkStackMXvsAtomics(b *testing.B) {

	const stackSize = 500
	b.Run("stackAtomics", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			wg.Add(stackSize)
			stck := exercises.NewLfStack[int]()
			for i := range stackSize {
				go func(i int) {
					defer wg.Done()
					stck.Push(i)
				}(i)
			}
			wg.Wait()
		}

	})
	b.Run("stackMx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			stck := exercises.NewStackMx[int]()
			var wg sync.WaitGroup
			wg.Add(stackSize)
			for i := range stackSize {
				go func(i int) {
					defer wg.Done()
					stck.Push(i)
				}(i)
			}
			wg.Wait()
		}
	})

}
