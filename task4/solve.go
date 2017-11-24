package main

import "fmt"
import "unicode"

func RemoveEven(inp []int) []int {
	out := make([]int, len(inp))
	j := 0
	for i := 0; i < len(inp); i++ {
		if inp[i]%2 != 0 {
			out[j] = inp[i]
			j++
		}
	}
	out = out[:j]
	return out
}

func PowerGenerator(a int) func() int {
	base := 1
	pow := a
	return func() int {
		base = base * pow
		return base
	}
}

func DifferentWordsCount(str string) (d uint) {
	d = 0
	arr := make([]string, 0)
	var s string
	var k int
	for i := 0; i < len(str); i++ {
		for ; i < len(str) && !unicode.IsLetter(rune(str[i])); i++ {
		}
		if i < len(str) {
			s = s + string(unicode.ToLower(rune(str[i])))
			i++
			for ; i < len(str) && unicode.IsLetter(rune(str[i])); i++ {
				s = s + string(unicode.ToLower(rune(str[i])))
			}
			for k = 0; k < len(arr) && !(arr[k] == s); k++ {
			}
			if k == len(arr) {
				arr = append(arr, s)
				d++
			}
			s = ""
		}
	}
	return
}
func main() {
	fmt.Println(DifferentWordsCount("!!!1!!;;'"))
}
