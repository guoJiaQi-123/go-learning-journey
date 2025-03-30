package main

import "fmt"

// 常量还可以用作枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

// iota
const (
	CategoryBooks    = iota // 0
	CategoryHealth          // 1
	CategoryClothing        // 2
)

const (
	IgEggsAllergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                // 1 << 1 which is 00000010
	IgNuts                     // 1 << 2 which is 00000100
	IgStrawberries             // 1 << 3 which is 00001000
	IgShellfish                // 1 << 4 which is 00010000
)

const (
	_          = iota             // ignore first value by assigning to blank identifier
	KBByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                            // 1 << (10*2)
	GB                            // 1 << (10*3)
	TB                            // 1 << (10*4)
	PB                            // 1 << (10*5)
	EB                            // 1 << (10*6)
	ZB                            // 1 << (10*7)
	YB                            // 1 << (10*8)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	println(a, b, c)
}
