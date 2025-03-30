package main

import "fmt"

var (
	e int    = 500
	f string = "多变量声明"
)

func main() {
	// 多变量声明

	// 方法一
	var a, b int = 100, 200
	var g, h = 600, "多变量"
	fmt.Println("a =", a, ",b =", b)
	fmt.Println("g =", g, ",h =", h)

	// 方法二：多用于全局变量的声明
	var (
		c int = 300
		d int = 400
	)

	fmt.Println(c)
	fmt.Println(d)
}
