package main

import (
	"fmt"
	"sync"
	"time"
)

/*
   Cond は効率的に条件を満たす Wait を実現できる。（ for と sleep の組み合わせを使わなくてよい）
   Mutex に対して Wait 機能を付加したもの。 Signal 、 Broadcast で Wait 中の Cond のブロックを解除できる。
*/
func main() {
	m := &sync.Mutex{}
	cond := sync.NewCond(m)

	go func() {
		cond.L.Lock() // L に Mutex が入っている
		defer cond.L.Unlock()
		fmt.Println("Start waiting.")
		cond.Wait()
		fmt.Println("End waiting.")
	}()

	fmt.Println("Exec someting.")
	time.Sleep(1 * time.Second) // 何か処理
	cond.Signal()               // これが無いと、上記の goroutine が再開できないので、先に メイン goroutine が終了する。
	time.Sleep(1 * time.Second) // メイン goroutine 終了の抑止
}
