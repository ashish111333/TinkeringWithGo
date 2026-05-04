package concurrency

import (
	"strconv"
	"sync"
	"testing"
)

func TestShardedKVSetGetDelete(t *testing.T) {
	kv := NewShardedKV(8)

	kv.Set("player:1", 42)
	if got, ok := kv.Get("player:1"); !ok || got != 42 {
		t.Fatalf("Get() = (%d, %v), want (42, true)", got, ok)
	}

	kv.Delete("player:1")
	if _, ok := kv.Get("player:1"); ok {
		t.Fatalf("expected key to be deleted")
	}
}

func TestShardedKVLen(t *testing.T) {
	kv := NewShardedKV(4)

	kv.Set("a", 1)
	kv.Set("b", 2)
	kv.Set("c", 3)

	if got, want := kv.Len(), 3; got != want {
		t.Fatalf("Len() = %d, want %d", got, want)
	}
}

func TestShardedKVConcurrentIncrement(t *testing.T) {
	kv := NewShardedKV(16)
	const goroutines = 100
	const incPerG = 1000
	key := "global_counter"

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()

			for range incPerG {
				kv.Increment(key, 1)
			}
		}()
	}
	wg.Wait()

	if got, ok := kv.Get(key); !ok || got != goroutines*incPerG {
		t.Fatalf("Get(%q) = (%d, %v), want (%d, true)", key, got, ok, goroutines*incPerG)
	}
}

func TestShardedKVConcurrentSetGet(t *testing.T) {
	kv := NewShardedKV(32)
	const keys = 1000

	var wg sync.WaitGroup
	wg.Add(keys)
	for i := range keys {
		go func(i int) {
			defer wg.Done()
			k := "k:" + strconv.Itoa(i)
			kv.Set(k, int64(i))
		}(i)
	}
	wg.Wait()

	if got, want := kv.Len(), keys; got != want {
		t.Fatalf("Len() = %d, want %d", got, want)
	}

	for i := 0; i < keys; i++ {
		k := "k:" + strconv.Itoa(i)
		if got, ok := kv.Get(k); !ok || got != int64(i) {
			t.Fatalf("Get(%q) = (%d, %v), want (%d, true)", k, got, ok, i)
		}
	}
}
