// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
// 結構體是一種複合的數據類型，用來組合多種不同（或相同）的數據類型到一個單一的類型中。
// type 關鍵字表示我們要定義一個新的類型，person 是這個新類型的名稱，而 struct 告訴 Go，這個新類型將是一個結構體。
// name string: 這表示 person 結構體有一個名為 name 的字段，其數據類型為 string。
// age int: 這表示 person 結構體有一個名為 age 的字段，其數據類型為 int
type person struct {
	name string
	age  int
}

// `newPerson` constructs a new person struct with the given name.
func newPerson(name string) *person {
	// You can safely return a pointer to local variable
	// as a local variable will survive the scope of the function.
	// 創建 person 的實例（或稱為對象），並給它賦值，
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An `&` prefix yields a pointer to the struct.
	// &: 此符號是取地址運算符。返回該變量的內存地址
	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the
	// pointers are automatically dereferenced.
	// 在 Go 語言中，當你有一個指針指向某個結構時，你可以直接使用 . 來訪問其欄位，而不需要明確的解參考。
	sp := &s
	fmt.Println(sp.age) // Output: 50

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age) // Output: 51

	// If a struct type is only used for a single value, we don't
	// have to give it a name. The value can have an anonymous
	// struct type. This technique is commonly used for
	// [table-driven tests](testing-and-benchmarking).
	// 定義了一個匿名結構體，並且初始化匿名結構體
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog) // Output: {Rex true}
}
