package handler

import (
	"context"
	"fmt"
	"time"
)

func WaitCancel() {
	// # キャンセル通知があるまで処理を待機する場合、context.Context.Doneメソッドで得られるチャネルからの通知を待つ
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	go func() { fmt.Println("another goroutine") }()
	fmt.Println("STOP")
	<-ctx.Done()
	fmt.Println("start moving")
}
