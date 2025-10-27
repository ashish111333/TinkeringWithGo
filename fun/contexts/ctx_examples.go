package contexts

import (
	"context"
	"fmt"
	"time"
)

func SimulateIoWork(ctx context.Context) error {

	fmt.Println("starting io work...")
	chn := ctx.Done()
	for {
		select {
		case <-chn:
			fmt.Println("stopping work")
			return ctx.Err()
		default:
			time.Sleep(time.Millisecond * 3000)
		}

	}

}
