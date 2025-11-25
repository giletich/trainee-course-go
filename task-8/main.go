package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	
	"task-8/cake"
	cakepipeline "task-8/cake_pipeline"
	"task-8/printcakes"
)

const (
	K = 10000
	N = 100
	M = 50
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	numbers := make(chan int, K)
	cakes := make(chan cake.Cake, K)
	pkg := make(chan cake.Cake, K)

	cakepipeline.RunPipeline(K, M, N, numbers, cakes, pkg, ctx)

	printcakes.PrintCakes(pkg)
}
