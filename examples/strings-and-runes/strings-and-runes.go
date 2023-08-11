// A Go string is a read-only slice of bytes. The language
// and the standard library treat strings specially - as
// containers of text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// In other languages, strings are made of "characters".
// In Go, the concept of a character is called a `rune` - it's
// an integer that represents a Unicode code point.
// [This Go blog post](https://go.dev/blog/strings) is a good
// introduction to the topic.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` is a `string` assigned a literal value
	// representing the word "hello" in the Thai
	// language. Go string literals are UTF-8
	// encoded text.
	const s = "สวัสดี"

	// Since strings are equivalent to `[]byte`, this
	// will produce the length of the raw bytes stored within.
	fmt.Println("Len:", len(s))

	// Indexing into a string produces the raw byte values at
	// each index. This loop generates the hex values of all
	// the bytes that constitute the code points in `s`.
	// %x 是一個格式化輸出的動詞，它表示將輸入的整數轉換為十六進位表示。
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// To count how many _runes_ are in a string, we can use
	// the `utf8` package. Note that the run-time of
	// `RuneCountInString` depends on the size of the string,
	// because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by multiple UTF-8
	// code points, so the result of this count may be surprising.
	// utf8.RuneCountInString 是一個用來計算字串中 Unicode 字元數量的函式。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A `range` loop handles strings specially and decodes
	// each `rune` along with its offset in the string.
	// idx 是迴圈變數，表示字元在字串中的索引位置
	// runeValue 則是當前迴圈中的 Unicode 字元。
	for idx, runeValue := range s {
		// %#U 是一個格式化輸出的動詞，它表示將輸入的整數（這裡是 Unicode 碼位）轉換成其 Unicode 表示形式，並在前面加上 U+。
		// %d 則是另一個格式化動詞，用來輸出索引位置 idx
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the
	// `utf8.DecodeRuneInString` function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		fmt.Println("s[i:]: ", s[i:]);
		// utf8.DecodeRuneInString 解碼字串中的 UTF-8 編碼的『單個』 Unicode 字元（也稱為 Rune）的函式。
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// This demonstrates passing a `rune` value to a function.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// Values enclosed in single quotes are _rune literals_. We
	// can compare a `rune` value to a rune literal directly.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
