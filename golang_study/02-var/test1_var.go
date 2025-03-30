package main

import "fmt"

func main() {
	// 声明变量的四种方式
	// 方法一,var 变量名 变量类型 「变量值为默认值」
	var a int
	fmt.Println(a)

	// 方法二：var 变量名 变量类型 = 变量值
	var b int = 100
	fmt.Println(b)

	// 方法三：var 变量名 = 变量值
	var c = 200
	fmt.Println(c)

	// 方法四 变量名:=变量值
	d := 300
	fmt.Println(d)
}
