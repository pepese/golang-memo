package main

import (
	"fmt"
	"sync"
	"time"
)

/*
   Cond は効率的に条件を満たす Wait を実現できる。（ for と sleep の組み合わせを使わなくてよい）
*/
func main() {
	cond := sync.NewCond(&sync.Mutex{})

	cond.L.Lock()
	go func() {
		defer cond.L.Unlock()
		fmt.Println("Start waiting.")
		cond.Wait()
		fmt.Println("End waiting.")
	}()

	fmt.Println("Exec someting.")
	time.Sleep(1 * time.Second) // 何か処理
	cond.Signal()               // これが無いと、上記の goroutine が再開できないので、先に メイン goroutine が終了する。
	time.Sleep(1 * time.Second) // 何か処理
}
