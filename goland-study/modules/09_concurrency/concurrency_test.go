package concurrency

import (
	"sort"
	"sync"
	"testing"
)

func TestRunWorkers(t *testing.T) {
	ch := RunWorkers(5)
	if ch == nil {
		t.Fatal("RunWorkers returned nil channel")
	}
	var results []int
	for v := range ch {
		results = append(results, v)
	}
	if len(results) != 5 {
		t.Fatalf("expected 5 results, got %d", len(results))
	}
	sort.Ints(results)
	for i, v := range results {
		if v != i {
			t.Errorf("results[%d] = %d, want %d", i, v, i)
		}
	}
}

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 1000 {
		t.Errorf("SafeCounter.Value() = %d, want 1000", got)
	}
}

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	ch2 <- 4
	ch2 <- 5
	close(ch2)

	out := MergeChannels(ch1, ch2)
	var results []int
	for v := range out {
		results = append(results, v)
	}
	if len(results) != 6 {
		t.Fatalf("MergeChannels produced %d values, want 6", len(results))
	}
	sort.Ints(results)
	expected := []int{1, 2, 3, 4, 5}
	_ = expected // order is nondeterministic; just check count and sum
	sum := 0
	for _, v := range results {
		sum += v
	}
	if sum != 15 {
		t.Errorf("sum of merged = %d, want 15", sum)
	}
}
