package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, cancel1 := context.WithCancel(context.Background())

	daemonCtx, childCtx1 := asyncDo1(ctx1, cancel1)

	<-daemonCtx.Done()
	fmt.Println("daemonCtx is done")


	<-ctx1.Done()
	fmt.Println("ctx1 is done")

	<-childCtx1.Done()
	fmt.Println("childCtx1 is done via ctx1 cancel()")

}

func asyncDo1(ctx context.Context, cancelFunc context.CancelFunc) (daemonCtx,childCtx context.Context) {
	fmt.Println("doing ctx1")

	ctx2, cancelFunc2 := context.WithCancel(context.Background())
	childCtx, _ = context.WithTimeout(ctx,5*time.Second)
	go func(cancelFunc2 context.CancelFunc) {
		fmt.Println("start daemon...")
		time.Sleep(1 * time.Second)
		fmt.Println("daemon finished")
		defer cancelFunc2()
	}(cancelFunc2)

	defer cancelFunc()

	return ctx2,childCtx
}
