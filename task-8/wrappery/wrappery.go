package wrappery

import (
	"context"
	"sync"
	"task-8/cake"
	"time"
)

func Wrappery(cake, out chan cake.Cake, K int, M int, ctx context.Context) {
	wg := sync.WaitGroup{}

	for i := 1; i <= M; i++ {
		wg.Add(1)
		go WrapACake(i, cake, out, &wg, ctx)
	}

	wg.Wait()
}

func WrapACake(number int, in, out chan cake.Cake, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-in:
			if !ok {
				return
			}
			time.Sleep(time.Millisecond * time.Duration(number))

			select {
			case <-ctx.Done():
				return
			case out <- cake.Cake{
				CakeNumber: val.CakeNumber,
				BakedBy:    val.BakedBy,
				BakeTime:   val.BakeTime,
				PackedBy:   number,
				PackTime:   time.Now(),
			}:
			}
		}
	}
}