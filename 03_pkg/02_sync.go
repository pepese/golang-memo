/*
http://golang.jp/pkg/sync

func (m *Mutex) Lock()
Lockは、mをロックする。すでにロックされていたときは、呼び出したゴルーチンは「ロックが可能になるまでブロック」される。

RWMutexは、Reader/Writerの相互排他ロックで、ロック状態を「複数」のReaderまたは「ひとつ」のWriterが持つことができます。
*/
package main

import (
	"sync"
	"time"
)

var mu sync.RWMutex
var data map[string]string

func main() {
	data = map[string]string{"hoge": "fuga"}
	mu = sync.RWMutex{}
	go read()
	go read()
	go write()
	go read()
	time.Sleep(5 * time.Second)
}

func read() {
	println("read_start")
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(1 * time.Second)
	println("read_complete", data["hoge"])
}

func write() {
	println("write_start")
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(2 * time.Second)
	data["hoge"] = "piyo"
	println("write_complete")
}
