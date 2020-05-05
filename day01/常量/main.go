package main

import (
	"fmt"
	"math"
)

// 可导出的MAX常量
const (
	s1  = "golang" // 私有的
	MAX = 10       // 公有的
)

func constDemo() {
	const s string = "hello"
	const a, b = 3, 4
	fmt.Println(s, a, b)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
	const (
		s1  = "golang" // 私有的
		MAX = 10       // 公有的
	)
	fmt.Println(s1, MAX)
}
func main() {
	constDemo()
}
