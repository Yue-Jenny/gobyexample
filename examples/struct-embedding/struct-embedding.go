// Go supports _embedding_ of structs and interfaces
// to express a more seamless _composition_ of types.
// This is not to be confused with [`//go:embed`](embed-directive) which is
// a go directive introduced in Go version 1.16+ to embed
// files and folders into the application binary.

package main

import "fmt"

type base struct {
	num int
}

// 為 base 結構定義了一個方法 describe()。
// 方法在其名稱之前有一個特殊的參數（在此稱為接收器），該參數表示該方法作用的具體類型。
// 在這個例子中，describe() 的接收器是 (b base)，這意味著 describe() 是 base 類型的方法。
func (b base) describe() string {
	// %v 是一個占位符，它將被後面參數的值替換。
	return fmt.Sprintf("base with num=%v", b.num)
}

// A `container` _embeds_ a `base`. An embedding looks
// like a field without a name.
// 定義 container 結構體並嵌入 base
type container struct {
	base
	str string
}

func main() {

	// When creating structs with literals, we have to
	// initialize the embedding explicitly; here the
	// embedded type serves as the field name.
	// 創建一個 container 的實例，並明確初始化其嵌入的 base。
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base's fields directly on `co`,
	// e.g. `co.num`.
	// 直接從 container 的實例 co 訪問 base 的 num 字段，而不需要使用中間的 base 字段名。
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternatively, we can spell out the full path using
	// the embedded type name.
	fmt.Println("also num:", co.base.num)

	// Since `container` embeds `base`, the methods of
	// `base` also become methods of a `container`. Here
	// we invoke a method that was embedded from `base`
	// directly on `co`.
	// 由於 container 嵌入了 base，所以 base 的方法也變成了 container 的方法。
	// 因此，可以直接在 co 上調用 describe 方法，這實際上是從 base 嵌入的。
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Embedding structs with methods may be used to bestow
	// interface implementations onto other structs. Here
	// we see that a `container` now implements the
	// `describer` interface because it embeds `base`.
	// 由於 container 嵌入了具有 describe 方法的 base，container 也被視為實現了 describer 介面。
	// 可以通過嵌入來重用或組合功能，而不需要明確地在每個結構上實現方法。
	var d describer = co
	fmt.Println("describer:", d.describe())
}
