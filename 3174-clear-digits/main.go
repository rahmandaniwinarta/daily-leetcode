package main

import (
	"fmt"
)

func main() {
	fmt.Println(ClearDigits("abc"))
	fmt.Println(ClearDigits("cb34"))

}

func ClearDigits(s string) string {
	slice := []rune{}
	for _, huruf := range s {
		if huruf >= '0' && huruf <= '9' {
			if len(slice) > 0 {
				slice = slice[:len(slice)-1]
			}
		} else {
			slice = append(slice, huruf)
		}
	}
	return string(slice)
}
