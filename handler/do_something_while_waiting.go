package handler

import (
	"context"
	"fmt"
	"time"
)

// #キャンセル通知が来るまでタスクの受信・処理・待機するような実装
// select文は操作を多重化できる制御構文
// Doneメソッドで得られるチャネルからメッセージを受信するまで、taskチャネルからの作業を待機・処理する無限ループを繰り返す処理
func DoSomeThingWhileWaiting() {
	ctx, cancel := context.WithCancel(context.Background())
	task := make(chan int)
	go func() {
		for {
			select {
			// チャネル対する受信操作を記載する
			// いずれかのcaseを満たして処理を進められるまで待機する
			case <-ctx.Done():
				return
			case i := <-task:
				fmt.Println("get", i)
			default:
				fmt.Println("キャンセルされていない")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)

	for i := 0; 5 > i; i++ {
		task <- i
	}

	cancel()
}
