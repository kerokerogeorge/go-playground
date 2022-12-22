package handler

import (
	"context"
	"fmt"
	"time"
)

func WaitCancel() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() { fmt.Println("another goroutine") }()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("start moving")
}
