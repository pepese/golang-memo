package main

import (
	"fmt"
	"sync"
)

/*
	sync.WaitGroup は wg.Add で待ち数をカウントアップし、 wg.Done でカウントダウンされる。
	カウントが 0 になるまで wg.Wait() で待ちが発生する。
*/
func main() {
	var wg sync.WaitGroup // wg の作成、 new いらん

	wg.Add(1) // カウントアップ
	go func() {
		defer wg.Done() // カウントダウン
		fmt.Println("1st go-routing.")
	}()

	wg.Add(1) // カウントアップ
	go func() {
		defer wg.Done() // カウントダウン
		fmt.Println("2nd go-routing.")
	}()

	wg.Wait() // ブロック
	fmt.Println("complete.")
}
