package main

import (
	"fmt"
	"strconv"
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

func channelTask1(result chan string) {
	// 重い処理を想定
	time.Sleep(time.Second * 2)
	fmt.Println("channelTask finished!")

	result <- "channelTask result"
}

func channelTask2(num int, result chan string) {
	// 重い処理を想定
	time.Sleep(time.Second * time.Duration(num))
	fmt.Println("channelTask finished!")

	result <- strconv.Itoa(num) + "s channelTask result"
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
	go channelTask1(result)
	go task2()

	// resultの中に何も入ってなければ、入ってくるまで待つ仕様になっている
	fmt.Println(<-result)

	// goroutineが終わる前に main 関数自体が終わる為、待ち時間をつける。
	time.Sleep(time.Second * 3)

	/*
		複数のチャネルを待機する場合に使用 select
	*/
	// 最初に終わった処理を受ける
	result1 := make(chan string)
	result2 := make(chan string)
	go channelTask2(1, result1)
	go channelTask2(2, result2)
	select {
	case v1 := <-result1:
		fmt.Println(v1)
	case v2 := <-result2:
		fmt.Println(v2)
	}
}
