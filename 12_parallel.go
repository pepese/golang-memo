package main

import (
	"fmt"
	"time"
)

func task1() {
	// 重い処理を想定
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finished!")

}

func task2() {
	fmt.Println("task2 finished!")
}

func channelTask(result chan string) {
	// 重い処理を想定
	time.Sleep(time.Second * 2)
	fmt.Println("channelTask finished!")

	result <- "channelTask result"
}

// メイン関数
func main() {
	/*
		goroutine
		return を受け取れない並列処理
	*/
	// goを付けて並行処理にする
	go task1()
	go task2()
	// goroutineが終わる前に main 関数自体が終わる為、待ち時間をつける。
	time.Sleep(time.Second * 3)

	/*
		channel
		return を受け取れる並列処理
	*/
	// chan型で受け渡すデータの型はstring
	result := make(chan string)

	// goを付けて並行処理にする
	go channelTask(result)
	go task2()

	// resultの中に何も入ってなければ、入ってくるまで待つ仕様になっている
	fmt.Println(<-result)

	// goroutineが終わる前に main 関数自体が終わる為、待ち時間をつける。
	time.Sleep(time.Second * 3)
}
