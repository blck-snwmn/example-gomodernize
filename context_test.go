package main

import (
	"context"
	"testing"
	"time"
)

// TestUsingContextWithCancel demonstrates a test using context.WithCancel
// This can be modernized to use t.Context() in Go 1.24
func TestUsingContextWithCancel(t *testing.T) {
	// Pre Go 1.24 approach: manually create a context with timeout/cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run a function that uses the context
	result := someAsyncFunction(ctx)

	// Check the result
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

// someAsyncFunction simulates a function that uses a context
func someAsyncFunction(ctx context.Context) int {
	// Simulate some work
	select {
	case <-time.After(100 * time.Millisecond):
		return 42
	case <-ctx.Done():
		return 0
	}
}
