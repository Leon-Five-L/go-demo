package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func stopChannel(ctx context.Context) {
	c := make(chan bool)

	defer func() {
		if p := recover(); p != nil {
			fmt.Println("stopping grpc server panic", "panic", p)
		}
		fmt.Println("end")
	}()

	go func() {
		doSomethingforChan()
		close(c)
	}()

	select {
	case <-c:
		fmt.Println("stop grpc server")
	case <-ctx.Done():
		fmt.Println("noting happened")
	}
}

func doSomethingforChan() {
	fmt.Println("do something")
}

func TestChannel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go stopChannel(ctx)
	time.Sleep(3 * time.Second)
	cancel()
}
