package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"task-8/cake"
	cakepipeline "task-8/cake_pipeline"
	"task-8/printcakes"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	K := 10000
	N := 100
	M := 50

	numbers := make(chan int, K)
	cakes := make(chan cake.Cake, K)
	pkg := make(chan cake.Cake, K)

	wg.Add(1)
	go func() {
		defer wg.Done()
		cakepipeline.RunPipeline(K, M, N, numbers, cakes, pkg, ctx)
	}()

	select {
	case <-sigCh:
		cancel()
	case <-cakepipeline.WaitForPipeline(&wg):

	}

	printcakes.PrintCakes(pkg)
}

