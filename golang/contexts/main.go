package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	// context are immutable
	ctx = context.WithValue(ctx, "username", "Gbenga")

	// add ability to cancel the context
	// ctx, cancelCtx := context.WithCancel(ctx)

	// add a deadline cancellation to the context
	// withDeadline cancels the the context automatically after the time elapsed
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)

	// If we don't really know the time but you know the duration after which a context should be canceled,
	// then use withTimeout passing in Time.Duration
	// ctx, cancelCtx := context.WithTimeout(ctx, time.Seconds * 2)

	// If the cancellation ever occurs because of a deadline, then it can be useful to call the cancelCtx function to clean up resources
	// created by the context just to be on a safer end. This is an idempotent call also.
	defer cancelCtx()

	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("The value of username is %s", ctx.Value("username"))
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}

			fmt.Println("doSomething finished")
			return

		case number := <-printCh:
			fmt.Printf("I read %d\n", number)
		}
	}
}

func main() {
	ctx := context.Background()

	doSomething(ctx)
}
