package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const delay = 100 * time.Millisecond

func do(ctx context.Context, id int) {
	select {
	case <-ctx.Done():
		break
	case <-time.After(delay):
		fmt.Printf("%d done\n", id)
		break
	}
}

func main() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)

	timeout := delay
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup
	for id := 0; id < 1000; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			do(ctx, id)
		}(id)
	}
	wg.Wait()
}
