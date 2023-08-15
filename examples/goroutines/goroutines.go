// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	// 使用 go 關鍵字呼叫 f 函式，創建一個 Goroutine 來執行它。
	// 表示 f("goroutine") 會在一個獨立的執行緒中執行，並且不會阻塞主執行緒，所以程式會繼續執行下一個語句，而不等待 f 函式完成。
	go f("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	// 另一個使用 Goroutine 的例子。它定義了一個匿名函式，並傳遞字串 "going" 作為參數。這個匿名函式會在獨立的 Goroutine 中執行，同樣地不會阻塞主執行緒。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now. Wait for them to finish
	// (for a more robust approach, use a [WaitGroup](waitgroups)).
	// 主執行緒暫停一秒鐘
	// 但在實際應用中，可能會使用更強大的同步機制，如 sync.WaitGroup。
	time.Sleep(time.Second)
	fmt.Println("done")
}
