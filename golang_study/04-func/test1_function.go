package main

import "fmt"

// Go 函数可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
	fmt.Println(" x ")
}
