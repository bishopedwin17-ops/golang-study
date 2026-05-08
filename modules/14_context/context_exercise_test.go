package contextexercise

import (
	"context"
	"testing"
	"time"
)

func TestSlowOperationCompletes(t *testing.T) {
	ctx := context.Background()
	err := SlowOperation(ctx, 10*time.Millisecond)
	if err != nil {
		t.Errorf("unexpected error for non-cancelled context: %v", err)
	}
}

func TestSlowOperationCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately

	err := SlowOperation(ctx, 1*time.Second)
	if err == nil {
		t.Error("expected error when context is cancelled, got nil")
	}
}

func TestFetchWithTimeoutExpires(t *testing.T) {
	// operation takes longer than the 100ms timeout
	err := FetchWithTimeout(500 * time.Millisecond)
	if err == nil {
		t.Error("expected timeout error, got nil")
	}
}

func TestFetchWithTimeoutSucceeds(t *testing.T) {
	// operation finishes within 100ms
	err := FetchWithTimeout(10 * time.Millisecond)
	if err != nil {
		t.Errorf("unexpected error for fast operation: %v", err)
	}
}
