package main

import (
	"fmt"
	"sync"
)

/*
   Mutex は、相互排他 mutual exclusion の略。
   クリティカルセッション（プログラムが共有リソースに対する排他アクセスを必要とする場所）をコントロールする。
   RWMutex もあり、読み込みと書き込みを意識した Lock （ Lock 、 RLock ）も可能。
   また、 sync.Locker は interface で Lock と Unlock を提供する。
*/
func main() {
	var count int // 共有リソース
	var lock sync.Mutex

	inc := func(wg *sync.WaitGroup) {
		defer wg.Done() // ちなみに defer は最後に記述された順に実行される

		lock.Lock()         // 自分が Lock を獲得できるまでブロックされる
		defer lock.Unlock() // Lock を解放する
		count++
		fmt.Println("Incremented: ", count)
	}

	dec := func(wg *sync.WaitGroup) {
		defer wg.Done()

		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Println("Incremented: ", count)
	}

	// サブ goroutine の Lock 中にメイン goroutine が完了することを防ぐ wg
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 5; i++ {
		go inc(&wg)
		go dec(&wg)
	}
	wg.Wait()
	fmt.Println("complete.")
}
