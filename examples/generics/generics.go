// Starting with version 1.18, Go has added support for
// _generics_, also known as _type parameters_.
// generics（泛型）
package main

import "fmt"

// As an example of a generic function, `MapKeys` takes
// a map of any type and returns a slice of its keys.
// This function has two type parameters - `K` and `V`;
// `K` has the `comparable` _constraint_, meaning that
// we can compare values of this type with the `==` and
// `!=` operators. This is required for map keys in Go.
// `V` has the `any` constraint, meaning that it's not
// restricted in any way (`any` is an alias for `interface{}`).
// 此函數接受一個鍵值對型態為 K 和 V 的 map（m map[K]V），並返回一個 K 型別的切片。
func MapKeys[K comparable, V any](m map[K]V) []K {
	// 創建了一個 K 型別的切片 r，它的初始長度為 0，容量被設定為 m 的長度。
	r := make([]K, 0, len(m))
	// 遍歷 map m 中的所有鍵。
	for k := range m {
		// 當前迭代到的鍵 k 添加到切片 r 中。
		// 用法如下：
		// slice = append(slice, elem1, elem2)
		// slice = append(slice, anotherSlice...)
		// 範例：
		// slice = append([]byte("hello "), "world"...)
		r = append(r, k)
	}
	return r
}

// As an example of a generic type, `List` is a
// singly-linked list with values of any type.
// 泛型參數 [T any]
// List 是一個泛型 singly-linked list 的定義
type List[T any] struct {
	// 指向 element[T] 結構體
	head, tail *element[T]
}

type element[T any] struct {
	// next: 指向下一個 element[T] 的指針
	next *element[T]
	// val: 是這個元素存儲的實際數據，其型別為 T。
	val T
}

// We can define methods on generic types just like we
// do on regular types, but we have to keep the type
// parameters in place. The type is `List[T]`, not `List`.
// List[T] 泛型鏈表結構體的 Push 方法
func (lst *List[T]) Push(v T) {
	// 檢查鏈表的尾部（tail）是否為空，也即檢查鏈表是否是空的。
	if lst.tail == nil {
		// 如果鏈表是空的，創建一個新的元素並賦值給 head，這表示新元素是鏈表的第一個元素。
		lst.head = &element[T]{val: v}
		// 這是鏈表的唯一元素，因此它既是頭部也是尾部。
		lst.tail = lst.head
	} else {
		// 在尾部後面創建一個新的元素。
		lst.tail.next = &element[T]{val: v}
		// 更新尾部指針，使其指向新添加的元素。
		lst.tail = lst.tail.next
	}
}

// List[T] 泛型鏈表結構體的 GetAll 方法
func (lst *List[T]) GetAll() []T {
	// 初始化一個空的 T 型別陣列
	var elems []T
	// 遍歷鏈表，從頭部開始，並在每次迭代中移動到下一個元素，直到達到鏈表的尾部。
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// When invoking generic functions, we can often rely
	// on _type inference_. Note that we don't have to
	// specify the types for `K` and `V` when
	// calling `MapKeys` - the compiler infers them
	// automatically.
	fmt.Println("keys:", MapKeys(m))

	// ... though we could also specify them explicitly.
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
