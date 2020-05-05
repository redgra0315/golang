package main

import "fmt"

// 算术运算符

func operatorDemo() {
	a := 21
	b := 10
	var c int

	c = a + b
	fmt.Printf("a + b = %d\n", c)

	a++
	fmt.Println(a)
	b--
	fmt.Println(b)
}

// 关系运算符
func relationDemo() {
	d := 21
	e := 10
	if d == e {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}
	if d > e {
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}
	if d < e {
		fmt.Println("a < b")
	} else {
		fmt.Println("a >= b")
	}
}

// 逻辑运算符
func logicDemo() {
	a := true
	b := false
	if a && b {
		fmt.Println("a和b 都是真")
	} else {
		fmt.Println("都是假")
	}
	fmt.Println("=============")
	if a || b {
		fmt.Println("a和b 都是至少一个真")
	} else {
		fmt.Println("都是假")
	}

	if !a {
		fmt.Println("a为假")
	} else {
		fmt.Println("a为真")
	}
}

// 位运算符
func byteComputeDemo() {
	var a uint = 60
	var b uint = 13
	var c uint
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)

	c = a & b
	fmt.Printf("a&b= %d\n", c)

	c = a | b
	fmt.Printf("a | b =%d\n", c)

	c = a ^ b
	fmt.Printf("a ^ b =%d\n", c)

	c = a << 2 // 60 * 4 = 240
	// a <<n 左移n位=乘以2的n次方
	fmt.Printf("a << 2=%d\n", c)

	c = a >> 2 // 60 * 4 = 15
	// a >> n 右移n位=乘以2的n次方
	a >>= 2
	fmt.Printf("a >>  2=%d\n", a)
}
func main() {
	//operatorDemo()
	//relationDemo()
	//logicDemo()
	byteComputeDemo()
}
