// Go supports _methods_ defined on struct types.
// Go 提供了定義在 struct 上的方法。
package main

import "fmt"

type rect struct {
	width, height int
}

// This `area` method has a _receiver type_ of `*rect`.
// r 是 rect 的指標，但當你使用 r.width 和 r.height 時，Go 語言允許你直接訪問結構的字段，而不需要明確地解參考。
// *rect 是指標類型的表示，*rect，它表示 "一個指向 rect 的指標"。
func (r *rect) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or value
// receiver types. Here's an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go automatically handles conversion between values
	// and pointers for method calls. You may want to use
	// a pointer receiver type to avoid copying on method
	// calls or to allow the method to mutate the
	// receiving struct.
	// 取址操作（&）： 使用 & 來取得一個變數的記憶體位址。
	// 當你有一個指針指向某個結構時，你可以直接使用 . 來訪問其欄位，而不需要明確的解參考。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
