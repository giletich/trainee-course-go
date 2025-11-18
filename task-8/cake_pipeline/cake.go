package cakepipeline

import (
	"context"
	"sync"
	"task-8/bakery"
	"task-8/cake"
	"task-8/wrappery"
)

func RunPipeline(K int, N int, M int, numbers chan int, cakes, pkg chan cake.Cake, ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(cakes)
		bakery.Bakery(numbers, cakes, K, N, ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(pkg)
		wrappery.Wrappery(cakes, pkg, K, M, ctx)
	}()

	wg.Wait()
}

func WaitForPipeline(wg *sync.WaitGroup) chan struct{} {
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}
