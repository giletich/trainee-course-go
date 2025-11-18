package bakery

import (
	"context"
	"sync"
	"task-8/cake"
	"time"
)

func Bakery(numbers chan int, cakes chan cake.Cake, K int, N int, ctx context.Context) {
	wg := sync.WaitGroup{}
	

	for i := 1; i <= N; i++ {
		wg.Add(1)
		go BakeACake(i, numbers, cakes, &wg, ctx)
	}

	for i := 1; i <= K; i++ {
		select {
		case <-ctx.Done():
			return
		case numbers <- i:
		}
	}
	close(numbers)

	wg.Wait()
}

func BakeACake(number int, in chan int, out chan cake.Cake, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-in:
			if !ok {
				return
			}
			time.Sleep(time.Duration(number) * time.Millisecond)

			select {
			case <-ctx.Done():
				return
			case out <- cake.Cake{
				CakeNumber: val,
				BakedBy:    number,
				BakeTime:   time.Now(),
			}:
			}
		}
	}
}