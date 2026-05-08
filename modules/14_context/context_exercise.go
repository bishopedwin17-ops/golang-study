package contextexercise

import (
	"context"
	"time"
)

// SlowOperation simulates a slow task that takes `duration` to complete.
// It must respect ctx cancellation — if ctx is cancelled before the operation
// finishes, return ctx.Err().
// TODO: Implement using select and time.After.
func SlowOperation(ctx context.Context, duration time.Duration) error {
	// Your code here
	return nil
}

// FetchWithTimeout wraps SlowOperation with a 100ms timeout.
// Create the context with timeout inside this function.
// TODO: Implement.
func FetchWithTimeout(duration time.Duration) error {
	// Your code here
	return nil
}

// keep
var _ = context.Background
