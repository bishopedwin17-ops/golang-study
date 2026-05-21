package concurrency

import "sync"

// RunWorkers launches `n` goroutines. Each goroutine sends its worker ID (0 to n-1)
// into the returned channel. The channel should be closed when all workers are done.
// Use sync.WaitGroup.
// TODO: Implement.
func RunWorkers(n int) <-chan int {
	// Your code here
	return nil
}

// SafeCounter is a goroutine-safe counter.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment increments the counter safely.
// TODO: Implement using mu.Lock() / defer mu.Unlock().
func (c *SafeCounter) Increment() {
	// Your code here
}

// Value returns the current counter value safely.
// TODO: Implement.
func (c *SafeCounter) Value() int {
	// Your code here
	return 0
}

// MergeChannels takes two read-only int channels and merges them into one.
// The output channel closes when BOTH inputs are exhausted.
// TODO: Implement using goroutines and sync.WaitGroup.
func MergeChannels(ch1, ch2 <-chan int) <-chan int {
	// Your code here
	return nil
}
