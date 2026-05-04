package concurrency

import (
	"strconv"
	"sync"
	"testing"
)

func BenchmarkKVStoresMixedWorkload(b *testing.B) {
	const (
		workers        = 32
		opsPerWorker   = 2000
		keyCardinality = 4096
	)

	keys := make([]string, keyCardinality)
	for i := range keys {
		keys[i] = "k:" + strconv.Itoa(i)
	}

	b.Run("simple_kv", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			kv := NewSimpleKV()
			runMixedOpsSimpleKV(kv, keys, workers, opsPerWorker)
		}
	})

	b.Run("sharded_kv_32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			kv := NewShardedKV(32)
			runMixedOpsShardedKV(kv, keys, workers, opsPerWorker)
		}
	})
}

func runMixedOpsSimpleKV(kv *SimpleKV, keys []string, workers int, opsPerWorker int) {
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(workerID int) {
			defer wg.Done()
			base := workerID * opsPerWorker
			for j := 0; j < opsPerWorker; j++ {
				k := keys[(base+j)%len(keys)]

				switch j % 10 {
				case 0, 1:
					kv.Get(k)
				case 2:
					kv.Delete(k)
				default:
					kv.Increment(k, 1)
				}
			}
		}(w)
	}
	wg.Wait()
}

func runMixedOpsShardedKV(kv *ShardedKV, keys []string, workers int, opsPerWorker int) {
	var wg sync.WaitGroup
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func(workerID int) {
			defer wg.Done()
			base := workerID * opsPerWorker
			for j := 0; j < opsPerWorker; j++ {
				k := keys[(base+j)%len(keys)]

				switch j % 10 {
				case 0, 1:
					kv.Get(k)
				case 2:
					kv.Delete(k)
				default:
					kv.Increment(k, 1)
				}
			}
		}(w)
	}
	wg.Wait()
}
