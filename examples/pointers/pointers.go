// Go supports <em><a href="https://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>,
// allowing you to pass references to values and records
// within your program.

package main

import "fmt"

// We'll show how pointers work in contrast to values with
// 2 functions: `zeroval` and `zeroptr`. `zeroval` has an
// `int` parameter, so arguments will be passed to it by
// value. `zeroval` will get a copy of `ival` distinct
// from the one in the calling function.
// 這個函式中對 ival 的修改不會影響 main 函式內的 i 變數，因為 ival 是按值傳遞的，所以在函式中對它的修改『只影響函式內部的複本』。
func zeroval(ival int) {
	ival = 0
	fmt.Println("ival:", ival);
}

// `zeroptr` in contrast has an `*int` parameter, meaning
// that it takes an `int` pointer. The `*iptr` code in the
// function body then _dereferences_ the pointer from its
// memory address to the current value at that address.
// Assigning a value to a dereferenced pointer changes the
// value at the referenced address.
// 接受一個整數指標（*int）作為參數 iptr
// 這裡的 *iptr 是對指標的「解參考」操作，它讓我們可以訪問和修改指向的變數。
// 這個函式的操作會對 main 函式內的 i 產生影響，因為我們修改的是指向同一個變數的指標。
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // Output: initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i) // Output: zeroval: 1

	// The `&i` syntax gives the memory address of `i`,
	// i.e. a pointer to `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // Output: zeroptr: 0

	// Pointers can be printed too.
	fmt.Println("pointer:", &i) // Output: pointer: 0x1400000e0c8
}
