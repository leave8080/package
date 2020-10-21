package main

import (
	"context"
	"fmt"
	"time"
)

func main1() {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	go test(ctx)
	time.Sleep(time.Second * 10)
}

func test(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ee")
			return ctx.Err()

		default:
			fmt.Println("de")
			time.Sleep(time.Second)
		}
	}

}
func main() {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()
	go test(ctx)
	//time.Sleep(time.Second * 10)
}
